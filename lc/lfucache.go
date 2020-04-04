package lc

import ()

type Entry struct {
	next  *Entry
	prev  *Entry
	key   int
	count int
	value int
}

func makeEntry(key, value int) Entry {
	var ent Entry
	ent.next = nil
	ent.prev = nil
	ent.key = key
	ent.value = value
	return ent
}

type LFUCache struct {
	capactiy int
	store    map[Entry]int
	head     *Entry
	tail     *Entry
}

func Constructor(capacity int) LFUCache {
	var cache LFUCache
	cache.capactiy = 0
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LFUCache) Get(key int) int {
	return -1
}

func (this *LFUCache) Put(key int, value int) {
}
