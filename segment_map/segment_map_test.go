package segment_map

import (
	"sync"
	"testing"
)
var SHARD_COUNT=32
type SegmentMap []*ConcurrentMapShard

type ConcurrentMapShard struct {
	cache map[string]interface{}
	lock sync.RWMutex
}

func (m SegmentMap) GetShard(key string) *ConcurrentMapShard {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}
func New() SegmentMap {
	m:=make(SegmentMap,SHARD_COUNT)
	for i:=0;i<SHARD_COUNT;i++{
		m[i]=&ConcurrentMapShard{
			cache: make(map[string]interface{}),
		}
	}
	return m
}
func fnv32(key string) uint32 {
	hash:=uint32(2166136261)
	const prime32 = uint32(16777619)
	for i:=0;i<len(key);i++{
		hash*=prime32
		hash^=uint32(key[i])
	}
	return hash
}
func (s SegmentMap) Get(key string) (interface{}, bool) {
	shard:=s.GetShard(key)
	shard.lock.RLock()
	defer shard.lock.RUnlock()
	val,ok:=shard.cache[key]
	return val,ok
}
func (s SegmentMap) Put(key string,value interface{}) {
	shard:=s.GetShard(key)
	shard.lock.Lock()
	defer shard.lock.Unlock()
	shard.cache[key]=value


}
func TestSegmentMap(t *testing.T) {

}

