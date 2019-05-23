package test

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

var ep = func(k, v interface{}) {
	fmt.Printf("key=%v, value=%v\n", k, v)
}

type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	return sm.data[k]
}

func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}

func NewSynchoriedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

func LockCount() {
	wg.Add(10)

	for i := 0; i < 5; i++ {
		go LockRead(i)
	}

	for i := 0; i < 5; i++ {
		go LockWrite(i)
	}

	wg.Wait()
}

func LockRead(i int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", i)

	v := count

	fmt.Printf("读goroutine %d 读取结束，值为:%d\n", i, v)

	wg.Done()
	rw.RUnlock()
}

func LockWrite(i int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n", i)

	v := rand.Intn(1000)

	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", i, v)

	wg.Done()
	rw.Unlock()
}

func SyncMap() {
	sm := NewSynchoriedMap()
	sm.Put("name", "chenyubo")
	sm.Put("sex", "man")
	sm.Put("company", "didi")

	sm.Each(ep)
}
