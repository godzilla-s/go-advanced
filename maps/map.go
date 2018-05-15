package maps

import (
	"sync"
)

// 简单的db
type db struct {
	mux   sync.Mutex
	cache map[string]interface{}
}

func newDB() *db {
	return &db{
		cache: make(map[string]interface{}),
	}
}

func (db *db) Add(key string, val interface{}) {
	db.mux.Lock()
	defer db.mux.Unlock()

	db.cache[key] = val
}

func (db *db) Del(key string) {
	db.mux.Lock()
	defer db.mux.Unlock()

	if _, ok := db.cache[key]; ok {
		delete(db.cache, key)
	}
}

func (db *db) Update(key string, newVal interface{}) {
	db.mux.Lock()
	defer db.mux.Unlock()

	if _, ok := db.cache[key]; ok {
		db.cache[key] = newVal
	}
}
