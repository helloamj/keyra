package engine

import (
	"hash/crc32"

	dataStore "github.com/helloamj/keyra/internal/datastore"
)

type Engine struct {
	shards    []*dataStore.DataStore
	numShards uint
}

func New(numShards uint) *Engine {
	shards := make([]*dataStore.DataStore, numShards)
	for i := uint(0); i < numShards; i++ {
		shards[i] = &dataStore.DataStore{}
	}
	return &Engine{
		shards:    shards,
		numShards: numShards,
	}
}

func (e *Engine) Set(key string, value []byte) {
	shard := e.getShard(key)
	shard.Set(key, value)
}

func (e *Engine) Get(key string) ([]byte, bool) {
	shard := e.getShard(key)
	val, ok := shard.Get(key)
	return val, ok
}

func (e *Engine) Delete(key string) {
	shard := e.getShard(key)
	shard.Delete(key)

}

func (e *Engine) getShard(key string) *dataStore.DataStore {
	hash := crc32.ChecksumIEEE([]byte(key))
	return e.shards[uint(hash)%e.numShards]
}
