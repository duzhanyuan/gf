package gmap

import (
	"sync"
)

type IntInt64Map struct {
	sync.RWMutex
	m map[int]int64
}

func NewIntInt64Map() *IntInt64Map {
	return &IntInt64Map{
        m: make(map[int]int64),
    }
}

// 哈希表克隆
func (this *IntInt64Map) Clone() *map[int]int64 {
	m := make(map[int]int64)
	this.RLock()
	for k, v := range this.m {
		m[k] = v
	}
    this.RUnlock()
	return &m
}

// 设置键值对
func (this *IntInt64Map) Set(key int, val int64) {
	this.Lock()
	this.m[key] = val
	this.Unlock()
}

// 批量设置键值对
func (this *IntInt64Map) BatchSet(m map[int]int64) {
	this.Lock()
	for k, v := range m {
		this.m[k] = v
	}
	this.Unlock()
}

// 获取键值
func (this *IntInt64Map) Get(key int) (int64) {
	this.RLock()
	val, _ := this.m[key]
	this.RUnlock()
	return val
}

// 删除键值对
func (this *IntInt64Map) Remove(key int) {
    this.Lock()
    delete(this.m, key)
    this.Unlock()
}

// 批量删除键值对
func (this *IntInt64Map) BatchRemove(keys []int) {
    this.Lock()
    for _, key := range keys {
        delete(this.m, key)
    }
    this.Unlock()
}

// 返回对应的键值，并删除该键值
func (this *IntInt64Map) GetAndRemove(key int) (int64) {
    this.Lock()
    val, exists := this.m[key]
    if exists {
        delete(this.m, key)
    }
    this.Unlock()
    return val
}

// 返回键列表
func (this *IntInt64Map) Keys() []int {
    this.RLock()
    keys := make([]int, 0)
    for key, _ := range this.m {
        keys = append(keys, key)
    }
    this.RUnlock()
    return keys
}

// 返回值列表(注意是随机排序)
func (this *IntInt64Map) Values() []int64 {
    this.RLock()
    vals := make([]int64, 0)
    for _, val := range this.m {
        vals = append(vals, val)
    }
    this.RUnlock()
    return vals
}

// 是否存在某个键
func (this *IntInt64Map) Contains(key int) bool {
    this.RLock()
    _, exists := this.m[key]
    this.RUnlock()
    return exists
}

// 哈希表大小
func (this *IntInt64Map) Size() int {
    this.RLock()
    len := len(this.m)
    this.RUnlock()
    return len
}

// 哈希表是否为空
func (this *IntInt64Map) IsEmpty() bool {
    this.RLock()
    empty := (len(this.m) == 0)
    this.RUnlock()
    return empty
}

// 清空哈希表
func (this *IntInt64Map) Clear() {
    this.Lock()
    this.m = make(map[int]int64)
    this.Unlock()
}
