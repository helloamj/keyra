package engine

import (
	"hash/crc32"

	dataStore "github.com/helloamj/keyra/internal/datastore"
)

type Engine struct {
	partition    []*dataStore.DataStore
	numPartition uint
}

func New(numPartition uint) *Engine {
	partition := make([]*dataStore.DataStore, numPartition)
	for i := uint(0); i < numPartition; i++ {
		partition[i] = &dataStore.DataStore{}
	}
	return &Engine{
		partition:    partition,
		numPartition: numPartition,
	}
}

func (e *Engine) Set(key string, value []byte) {
	partition := e.getPartition(key)
	partition.Set(key, value)
}

func (e *Engine) Get(key string) ([]byte, bool) {
	partition := e.getPartition(key)
	val, ok := partition.Get(key)
	return val, ok
}

func (e *Engine) Delete(key string) {
	partition := e.getPartition(key)
	partition.Delete(key)

}

func (e *Engine) getPartition(key string) *dataStore.DataStore {
	hash := crc32.ChecksumIEEE([]byte(key))
	return e.partition[uint(hash)%e.numPartition]
}
