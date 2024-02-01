package main

import (
	"context"
	"github.com/forestyc/playground/cmd/demo/redis-transfer/common"
	"github.com/forestyc/playground/cmd/demo/redis-transfer/proto/redispb"
	"github.com/forestyc/playground/pkg/concurrency/workpool"
	"github.com/forestyc/playground/pkg/utils"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"time"
)

type Import struct {
	file *utils.File
	wp   *workpool.WorkPool
}

func NewImport(file string) *Import {
	return &Import{
		wp:   workpool.NewWorkPool(),
		file: utils.NewFile(file),
	}
}

func (i *Import) Run() {
	i.wp.Start()
	defer i.wp.Stop()

	bytes, err := i.file.LoadFile()
	if err != nil {
		panic("load file failed," + err.Error())
	}
	var start uint32
	var header common.Header
	for {
		if start >= uint32(len(bytes)) {
			break
		}
		headerBytes := bytes[start : start+common.HeadLength]
		header.Unmarshal(headerBytes)
		start += common.HeadLength
		i.wp.AddJob(i.callback(header.Type, bytes[start:start+header.Length]))
		start += header.Length
	}
}

func (i *Import) callback(t uint32, message []byte) workpool.Job {
	return func(ctx context.Context) {
		switch t {
		case common.Hash:
			i.Hash(message)
		case common.String:
			i.String(message)
		case common.List:
			i.List(message)
		case common.Set:
			i.Set(message)
		case common.ZSet:
			i.ZSet(message)
		case common.Stream:
			i.Stream(message)
		default:
			zlog.Warn("Export failed", zap.Uint32("invalid key type", t))
		}
	}
}

func (i *Import) Hash(message []byte) {
	// value
	var pb redispb.Hash
	err := proto.Unmarshal(message, &pb)
	if err != nil {
		zlog.Error("Hash failed", zap.Error(err))
		return
	}
	for _, field := range pb.Fields {
		if err = r.HSet(context.Background(), pb.Key, field.Field, field.Value).Err(); err != nil {
			zlog.Error("Hash failed", zap.Error(err), zap.String("function", "r.HSet"),
				zap.String("key", pb.Key), zap.String("field", field.Field))
		}
	}
	// ttl
	if pb.Ttl > 0 {
		if err = r.Expire(context.Background(), pb.Key, time.Duration(pb.Ttl)).Err(); err != nil {
			zlog.Error("Hash failed", zap.Error(err), zap.String("function", "r.Expire"),
				zap.String("key", pb.Key))
		}
	}
	zlog.Info("Import key", zap.String("key", pb.Key))
}

func (i *Import) String(message []byte) {
	// value
	var pb redispb.String
	err := proto.Unmarshal(message, &pb)
	if err != nil {
		zlog.Error("String failed", zap.Error(err))
		return
	}
	if err = r.Set(context.Background(), pb.Key, pb.Value, time.Duration(pb.Ttl)).Err(); err != nil {
		zlog.Error("String failed", zap.Error(err), zap.String("function", "r.Set"),
			zap.String("key", pb.Key))
	}
	zlog.Info("Import key", zap.String("key", pb.Key))
}

func (i *Import) List(message []byte) {
	// value
	var pb redispb.List
	err := proto.Unmarshal(message, &pb)
	if err != nil {
		zlog.Error("List failed", zap.Error(err))
		return
	}
	for _, value := range pb.Values {
		if err = r.RPush(context.Background(), pb.Key, value).Err(); err != nil {
			zlog.Error("List failed", zap.Error(err), zap.String("function", "r.RPush"),
				zap.String("key", pb.Key))
		}
	}
	// ttl
	if pb.Ttl > 0 {
		if err = r.Expire(context.Background(), pb.Key, time.Duration(pb.Ttl)).Err(); err != nil {
			zlog.Error("List failed", zap.Error(err), zap.String("function", "r.Expire"),
				zap.String("key", pb.Key))
		}
	}
	zlog.Info("Import key", zap.String("key", pb.Key))
}

func (i *Import) Set(message []byte) {
	// value
	var pb redispb.Set
	err := proto.Unmarshal(message, &pb)
	if err != nil {
		zlog.Error("Set failed", zap.Error(err))
		return
	}
	for _, member := range pb.Members {
		if err = r.SAdd(context.Background(), pb.Key, member).Err(); err != nil {
			zlog.Error("Set failed", zap.Error(err), zap.String("function", "r.SAdd"),
				zap.String("key", pb.Key))
		}
	}
	// ttl
	if pb.Ttl > 0 {
		if err = r.Expire(context.Background(), pb.Key, time.Duration(pb.Ttl)).Err(); err != nil {
			zlog.Error("Set failed", zap.Error(err), zap.String("function", "r.Expire"),
				zap.String("key", pb.Key))
		}
	}
	zlog.Info("Import key", zap.String("key", pb.Key))
}

func (i *Import) ZSet(message []byte) {
	// value
	var pb redispb.ZSet
	err := proto.Unmarshal(message, &pb)
	if err != nil {
		zlog.Error("ZSet failed", zap.Error(err))
		return
	}
	for _, member := range pb.Members {
		if err = r.ZAdd(context.Background(), pb.Key, &redis.Z{
			Score:  member.Score,
			Member: member.Member,
		}).Err(); err != nil {
			zlog.Error("ZSet failed", zap.Error(err), zap.String("function", "r.ZAdd"),
				zap.String("key", pb.Key))
		}
	}
	// ttl
	if pb.Ttl > 0 {
		if err = r.Expire(context.Background(), pb.Key, time.Duration(pb.Ttl)).Err(); err != nil {
			zlog.Error("ZSet failed", zap.Error(err), zap.String("function", "r.Expire"),
				zap.String("key", pb.Key))
		}
	}
	zlog.Info("Import key", zap.String("key", pb.Key))
}

func (i *Import) Stream(message []byte) {
	// value
	var pb redispb.Stream
	err := proto.Unmarshal(message, &pb)
	if err != nil {
		zlog.Error("Stream failed", zap.Error(err))
		return
	}
	for _, m := range pb.Messages {
		var values []interface{}
		for idx := 0; idx < len(m.Field); idx++ {
			values = append(values, m.Field[idx], string(m.Value[idx]))
		}
		xargs := redis.XAddArgs{
			Stream: pb.Key,
			ID:     m.Id,
			Values: values,
		}
		if err = r.XAdd(context.Background(), &xargs).Err(); err != nil {
			zlog.Error("Stream failed", zap.Error(err), zap.String("function", "r.XAdd"),
				zap.String("key", pb.Key))
		}
	}
	// ttl
	if pb.Ttl > 0 {
		if err = r.Expire(context.Background(), pb.Key, redis.KeepTTL).Err(); err != nil {
			zlog.Error("Stream failed", zap.Error(err), zap.String("function", "r.Expire"),
				zap.String("key", pb.Key))
		}
	}
	zlog.Info("Import key", zap.String("key", pb.Key))
}
