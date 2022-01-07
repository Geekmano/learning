package rw_map

import (
	"sync"
	"testing"
)

type RWMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

func (m RWMap) GET(key interface{}) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	value, ok := m.m[key]
	return value, ok

}
func (m RWMap) Put(key, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m[key] = value
}
func (m RWMap) Del(key interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.m, key)

}

func TestRWMap(t *testing.T) {

}
