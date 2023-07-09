package main

import (
	"sync"
)

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4
}

type LRUCache struct {
	cap     int
	cache   map[int]int
	usedKey []int
	lock    *sync.Mutex
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:     capacity,
		cache:   make(map[int]int),
		usedKey: make([]int, 0),
		lock:    &sync.Mutex{},
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.get(key)
}

func (this *LRUCache) Put(key int, value int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if valueCached := this.get(key); valueCached == -1 {
		lenth := len(this.cache)
		if lenth >= this.cap {
			removeKey := this.usedKey[0]
			delete(this.cache, removeKey)
			tmp := make([]int, 0)
			tmp = append(tmp, this.usedKey[1:]...)
			this.usedKey = tmp
		}
	}
	this.recordKey(key)
	this.cache[key] = value
}

func (this *LRUCache) get(key int) int {
	value, ok := this.cache[key]
	if ok {
		this.recordKey(key)
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) recordKey(key int) {
	idx := -1
	for k, e := range this.usedKey {
		if e == key {
			idx = k
			break
		}
	}
	if idx != -1 {
		tmp := make([]int, 0)
		tmp = append(tmp, this.usedKey[:idx]...)
		tmp = append(tmp, this.usedKey[idx+1:]...)
		this.usedKey = tmp
	}
	this.usedKey = append(this.usedKey, key)

}
