package datastore

import (
	"sync"
)

type DataStore struct {
	data sync.Map
}

func newDataStore() *DataStore {
	return &DataStore{}
}

func (s *DataStore) Set(key string, value []byte) {
	s.data.Store(key, value)
}

func (s *DataStore) Get(key string) ([]byte, bool) {
	val, ok := s.data.Load(key)
	if !ok {
		return nil, false
	}
	return val.([]byte), true
}

func (s *DataStore) Delete(key string) {
	s.data.Delete(key)
}