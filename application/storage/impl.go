package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"lrucache/pkg/lrucache"
	"time"
)

type storageLRUImpl struct {
	cache lrucache.LRUCache[string, []byte]
}

type StorageLRUCfg struct {
	CacheSize int `arg:"--cache-size,env:CACHE_SIZE" default:"100"`
}

func NewStorageLRU(cfg StorageLRUCfg) Storage {
	return &storageLRUImpl{
		cache: lrucache.NewLRUCache[string, []byte](cfg.CacheSize),
	}
}

var _ Storage = (*storageLRUImpl)(nil)

func (s *storageLRUImpl) Get(ctx context.Context, key string, template interface{}) error {
	data, found := s.cache.Get(key)
	if !found {
		return fmt.Errorf("can`t get value from cache: %s", key)
	}

	if err := json.Unmarshal(data, template); err != nil {
		return fmt.Errorf("can`t unmarshal value: %w", err)
	}

	return nil
}

func (s *storageLRUImpl) Set(ctx context.Context, key string, value interface{}, dataTTL time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("can`t marshal value: %w", err)
	}

	added := s.cache.Add(key, data)

	if !added {
		return fmt.Errorf("can`t store value into cache: %s", key)
	}

	return nil
}
