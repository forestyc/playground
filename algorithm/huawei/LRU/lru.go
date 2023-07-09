package main

import (
	"sync"
)

func main() {
	lRUCache := Constructor(3)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	lRUCache.Put(3, 3)
	lRUCache.Put(4, 4)
	lRUCache.Get(4)
	lRUCache.Get(3)
	lRUCache.Get(2)
	lRUCache.Get(1)
	lRUCache.Put(5, 5)
	lRUCache.Get(1)
	lRUCache.Get(2)
	lRUCache.Get(3)
	lRUCache.Get(4)
	lRUCache.Get(5)
}

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	cap      int
	cache    map[int]*Node
	listHead *Node
	listTail *Node
	lock     *sync.Mutex
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		lock:  &sync.Mutex{},
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	node, ok := this.cache[key]
	if ok {
		this.moveToTial(node)
		return node.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	node, ok := this.cache[key]
	if !ok {
		if len(this.cache) >= this.cap {
			this.removeUnused()
		}
		node = &Node{}
		this.addNode(node)
	} else {
		this.moveToTial(node)
	}
	node.Key = key
	node.Value = value
	this.cache[key] = node
}

func (this *LRUCache) addNode(node *Node) {
	if this.listHead == nil {
		this.listHead = node
		this.listTail = node
	} else {
		if this.listHead.Next == nil {
			this.listHead.Next = node
		}
		// node变为尾结点
		node.Pre = this.listTail
		node.Next = nil
		this.listTail.Next = node
		// 记录尾结点
		this.listTail = node
	}
}

func (this *LRUCache) moveToTial(node *Node) {
	if len(this.cache) <= 1 {
		return
	}
	if node == this.listHead { // 目标为头结点，头结点后移
		this.listHead = node.Next
		this.listHead.Pre = nil
	} else if node != this.listTail { // 目标为中间节点，摘除目标节点
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	} else { // 目标为尾结点，return
		return
	}
	// node变为尾结点
	node.Pre = this.listTail
	node.Next = nil
	this.listTail.Next = node
	// 记录尾结点
	this.listTail = node
}

func (this *LRUCache) removeUnused() {
	delete(this.cache, this.listHead.Key)
	this.listHead = this.listHead.Next
}
