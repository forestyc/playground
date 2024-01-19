package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/forestyc/playground/pkg/redis"
	"log"
	"os"
	"path"
	"strings"
)

var (
	r                       *redis.Redis
	operation, keys, prefix string
)

func main() {
	var err error
	flag.StringVar(&operation, "operation", "", "import/export")
	flag.StringVar(&keys, "keys", "", "key1,key2,key3")
	flag.StringVar(&prefix, "prefix", "./", "output prefix")
	flag.Parse()
	r, err = redis.NewRedis(redis.Config{
		Address:          "140.143.163.171:6379",
		Password:         "k8s-node1#12345",
		MaxOpen:          1,
		IdleConns:        1,
		IdleTimout:       360,
		OperationTimeout: 60,
	})
	if err != nil {
		panic(err)
	}
	if operation == "export" {
		Export()
	} else {
		Import()
	}
}

func Export() {
	keys = strings.TrimSpace(keys)
	keys = strings.ReplaceAll(keys, " ", "")
	for _, key := range strings.Split(keys, ",") {
		dirname := path.Join(prefix, key)
		if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
			log.Println("run", err)
			return
		}
		switch keyType(key) {
		case "hash":
			dumpHash(key, dirname)
		case "string":
			dumpString(key, dirname)
		case "list":
			dumpList(key, dirname)
		case "set":
			dumpSet(key, dirname)
		case "zset":
			dumpZSet(key, dirname)
		case "stream":
		default:
			panic("invalid key type[" + key + "]")
		}
	}
}

func Import() {

}

func dumpHash(key, dirname string) {
	result := r.HKeys(context.Background(), key)
	if result.Err() != nil {
		log.Println("dumpHash", result.Err(), key)
		return
	}
	var fields []string
	result.ScanSlice(&fields)
	for _, field := range fields {
		// get value
		v := r.HGet(context.Background(), key, field)
		if v.Err() != nil {
			log.Println("dumpHash", v.Err(), key, field)
			continue
		}
		bytes, err := v.Bytes()
		if err != nil {
			log.Println("dumpHash", err, key, field)
			continue
		}
		// to file
		filename := path.Join(dirname, field)
		if err = toFile(filename, bytes); err != nil {
			log.Println("dumpHash", err, key, field)
			continue
		}
	}
}

func dumpString(key, dirname string) {
	// get value
	result := r.Get(context.Background(), key)
	if result.Err() != nil {
		log.Println("dumpString", result.Err(), key)
		return
	}
	// to file
	filename := path.Join(dirname, "key")
	if err := toFile(filename, []byte(result.Val())); err != nil {
		log.Println("dumpString", err, key)
		return
	}
}

func dumpList(key, dirname string) {
	// get value
	result := r.LRange(context.Background(), key, 0, -1)
	if result.Err() != nil {
		log.Println("dumpList", result.Err(), key)
		return
	}
	list := result.Val()
	bytes, err := json.Marshal(&list)
	if err != nil {
		log.Println("dumpZSet", result.Err(), key)
		return
	}
	// to file
	filename := path.Join(dirname, "key")
	if err = toFile(filename, bytes); err != nil {
		log.Println("dumpList", err, key)
		return
	}
}

func dumpSet(key, dirname string) {
	// get value
	result := r.SMembers(context.Background(), key)
	if result.Err() != nil {
		log.Println("dumpSet", result.Err(), key)
		return
	}
	set := result.Val()
	bytes, err := json.Marshal(&set)
	if err != nil {
		log.Println("dumpZSet", result.Err(), key)
		return
	}
	// to file
	filename := path.Join(dirname, "key")
	if err = toFile(filename, bytes); err != nil {
		log.Println("dumpSet", err, key)
		return
	}
}

func dumpZSet(key, dirname string) {
	// get value
	result := r.ZRangeWithScores(context.Background(), key, 0, -1)
	if result.Err() != nil {
		log.Println("dumpZSet", result.Err(), key)
		return
	}
	zset := result.Val()
	bytes, err := json.Marshal(&zset)
	if err != nil {
		log.Println("dumpZSet", result.Err(), key)
		return
	}
	// to file
	filename := path.Join(dirname, "key")
	if err = toFile(filename, bytes); err != nil {
		log.Println("dumpZSet", err, key)
		return
	}
}

func keyType(key string) string {
	result := r.Type(context.Background(), key)
	if result.Err() != nil {
		log.Println("exclusive", result.Err())
		return ""
	}
	return result.Val()
}

func toFile(filename string, buf []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(buf)
	return nil
}
