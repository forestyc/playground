package main

import (
	"context"
	"github.com/forestyc/playground/cmd/demo/redis-transfer/common"
	"github.com/forestyc/playground/cmd/demo/redis-transfer/proto/redispb"
	"github.com/forestyc/playground/pkg/concurrency/workpool"
	"github.com/forestyc/playground/pkg/utils"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"
)

type Export struct {
	data chan proto.Message
	keys []string
	wp   *workpool.WorkPool
	file *utils.File
}

func NewExport(keys, file string) *Export {
	os.Remove(file)
	keys = strings.TrimSpace(keys)
	keys = strings.ReplaceAll(keys, " ", "")
	return &Export{
		data: make(chan proto.Message, 1000),
		keys: strings.Split(keys, ","),
		wp:   workpool.NewWorkPool(),
		file: utils.NewFile(file),
	}
}

func (e *Export) Run() {
	e.dump()
	e.wp.Start()
	defer e.wp.Stop()
	for _, key := range e.keys {
		if strings.Contains(key, "*") {
			result := r.Keys(context.Background(), key)
			if result.Err() != nil {
				zlog.Error("Run failed", zap.Error(result.Err()), zap.String("function", "r.Keys"))
				continue
			}
			for _, k := range result.Val() {
				e.wp.AddJob(e.Callback(k))
			}
		} else {
			e.wp.AddJob(e.Callback(key))
		}

	}
}

func (e *Export) Close() {
	close(e.data)
}

func (e *Export) dump() {
	go func() {
		for data := range e.data {
			bytesContent, err := proto.Marshal(data)
			if err != nil {
				zlog.Error("dump failed", zap.Error(err), zap.String("function", "proto.Marshal"))
				continue
			}
			header := common.NewHeader(common.RedisType(data), uint32(len(bytesContent)))
			bytesHeader := header.Marshal()
			if err != nil {
				zlog.Error("Hash failed", zap.Error(err), zap.String("function", "proto.Marshal"))
				continue
			}
			// write header
			if err = e.file.AppendFile(bytesHeader); err != nil {
				zlog.Error("Hash failed", zap.Error(err), zap.String("function", "utils.AppendFile"))
				continue
			}
			// write content
			if err = e.file.AppendFile(bytesContent); err != nil {
				zlog.Error("Hash failed", zap.Error(err), zap.String("function", "utils.AppendFile"))
				continue
			}
		}
		e.file.Close()
		zlog.Info("dump complete")
	}()
}

func (e *Export) Callback(key string) workpool.Job {
	return func(ctx context.Context) {
		t := e.KeyType(key)
		switch t {
		case "hash":
			e.Hash(key)
		case "string":
			e.String(key)
		case "list":
			e.List(key)
		case "set":
			e.Set(key)
		case "zset":
			e.ZSet(key)
		case "stream":
			e.Stream(key)
		default:
			zlog.Warn("Export failed", zap.String("invalid key type", t))
		}
	}
}

func (e *Export) Hash(key string) {
	// ttl
	ttl, err := r.TTL(context.Background(), key).Result()
	if err != nil {
		zlog.Error("Hash failed", zap.Error(err), zap.String("function", "r.TTL"),
			zap.String("key", key))
		return
	}
	hash := redispb.Hash{
		Key: key,
		Ttl: int32(ttl),
	}
	var fields []string
	if err := r.HKeys(context.Background(), key).ScanSlice(&fields); err != nil {
		zlog.Error("Hash failed", zap.Error(err), zap.String("key", key))
		return
	}
	for _, field := range fields {
		// get value
		v, err := r.HGet(context.Background(), key, field).Bytes()
		if err != nil {
			zlog.Error("Hash failed", zap.Error(err), zap.String("function", "r.HGet"),
				zap.String("key", key), zap.String("field", field))
			continue
		}
		f := redispb.Field{
			Field: field,
			Value: v,
		}
		hash.Fields = append(hash.Fields, &f)
	}
	e.data <- &hash
	zlog.Info("Hash cached", zap.String("key", key))
}

func (e *Export) String(key string) {
	// get value
	v, err := r.Get(context.Background(), key).Bytes()
	if err != nil {
		zlog.Error("String failed", zap.Error(err), zap.String("function", "r.Get"),
			zap.String("key", key))
		return
	}
	// ttl
	ttl, err := r.TTL(context.Background(), key).Result()
	if err != nil {
		zlog.Error("String failed", zap.Error(err), zap.String("function", "r.TTL"),
			zap.String("key", key))
		return
	}
	e.data <- &redispb.String{
		Key:   key,
		Value: v,
		Ttl:   int32(ttl),
	}
	zlog.Info("String cached", zap.String("key", key))
}

func (e *Export) List(key string) {
	// ttl
	ttl, err := r.TTL(context.Background(), key).Result()
	if err != nil {
		zlog.Error("List failed", zap.Error(err), zap.String("function", "r.TTL"),
			zap.String("key", key))
		return
	}
	list := redispb.List{
		Key: key,
		Ttl: int32(ttl),
	}
	// get value
	var values [][]byte
	if err = r.LRange(context.Background(), key, 0, -1).ScanSlice(&values); err != nil {
		zlog.Error("List failed", zap.Error(err), zap.String("function", "r.LRange"),
			zap.String("key", key))
		return
	}
	for _, value := range values {
		list.Values = append(list.Values, value)
	}
	e.data <- &list
	zlog.Info("List cached", zap.String("key", key))
}

func (e *Export) Set(key string) {
	// ttl
	ttl, err := r.TTL(context.Background(), key).Result()
	if err != nil {
		zlog.Error("Set failed", zap.Error(err), zap.String("function", "r.TTL"),
			zap.String("key", key))
		return
	}
	set := redispb.Set{
		Key: key,
		Ttl: int32(ttl),
	}
	// get value
	members, err := r.SMembers(context.Background(), key).Result()
	if err != nil {
		zlog.Error("Set failed", zap.Error(err), zap.String("function", "r.SMembers"),
			zap.String("key", key))
		return
	}
	for _, m := range members {
		set.Members = append(set.Members, m)
	}
	e.data <- &set
	zlog.Info("Set cached", zap.String("key", key))
}

func (e *Export) ZSet(key string) {
	// ttl
	ttl, err := r.TTL(context.Background(), key).Result()
	if err != nil {
		zlog.Error("ZSet failed", zap.Error(err), zap.String("function", "r.TTL"),
			zap.String("key", key))
		return
	}
	zset := redispb.ZSet{
		Key: key,
		Ttl: int32(ttl),
	}
	// get value
	result := r.ZRangeWithScores(context.Background(), key, 0, -1)
	if result.Err() != nil {
		zlog.Error("ZSet failed", zap.Error(result.Err()), zap.String("function", "r.ZRangeWithScores"),
			zap.String("key", key))
		return
	}
	for _, value := range result.Val() {
		zm := redispb.ZMember{
			Member: value.Member.(string),
			Score:  value.Score,
		}
		zset.Members = append(zset.Members, &zm)
	}
	e.data <- &zset
	zlog.Info("ZSet cached", zap.String("key", key))
}

func (e *Export) Stream(key string) {
	// ttl
	ttl, err := r.TTL(context.Background(), key).Result()
	if err != nil {
		zlog.Error("Stream failed", zap.Error(err), zap.String("function", "r.TTL"),
			zap.String("key", key))
		return
	}
	stream := redispb.Stream{
		Key: key,
		Ttl: int32(ttl),
	}
	// get value
	xlenResult := r.XLen(context.Background(), key)
	if xlenResult.Err() != nil {
		zlog.Error("Stream failed", zap.Error(xlenResult.Err()), zap.String("function", "r.XLen"),
			zap.String("key", key))
		return
	}
	count, _ := xlenResult.Uint64()
	xreadResult := r.XRead(context.Background(), &goRedis.XReadArgs{
		Streams: []string{key, "0"},
		Count:   int64(count),
	})
	if xreadResult.Err() != nil {
		zlog.Error("Stream failed", zap.Error(xreadResult.Err()),
			zap.String("function", "r.XRead"),
			zap.String("key", key))
		return
	}
	for _, s := range xreadResult.Val() {
		for _, message := range s.Messages {
			xm := redispb.XMessage{Id: message.ID}
			for field, value := range message.Values {
				xm.Field = append(xm.Field, field)
				xm.Value = append(xm.Value, []byte(value.(string)))
			}
			stream.Messages = append(stream.Messages, &xm)
		}
	}
	e.data <- &stream
	zlog.Info("Stream cached", zap.String("key", key))
}

func (e *Export) KeyType(key string) string {
	result := r.Type(context.Background(), key)
	if result.Err() != nil {
		log.Println("exclusive", result.Err())
		return ""
	}
	return result.Val()
}
