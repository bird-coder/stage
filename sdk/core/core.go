/*
 * @Author: yujiajie
 * @Date: 2024-03-13 14:45:11
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 17:55:43
 * @FilePath: /stage/sdk/core/core.go
 * @Description:
 */
package core

import (
	"stage/sdk/storage/cache"
	"stage/sdk/storage/locker"
	"sync"

	"gorm.io/gorm"
)

var App Core = NewApp()

type Application struct {
	dbs     map[string]*gorm.DB
	configs map[string]any
	cache   cache.AdapterCache
	locker  locker.AdapterLocker

	mux sync.RWMutex
}

func NewApp() *Application {
	return &Application{
		dbs:     make(map[string]*gorm.DB),
		configs: make(map[string]any),
	}
}

func (e *Application) SetDb(key string, db *gorm.DB) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.dbs[key] = db
}

func (e *Application) GetDb(key string) *gorm.DB {
	e.mux.RLock()
	defer e.mux.RUnlock()
	return e.dbs[key]
}

func (e *Application) GetAllDb() map[string]*gorm.DB {
	return e.dbs
}

func (e *Application) SetConfig(key string, config any) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.configs[key] = config
}

func (e *Application) GetConfig(key string) any {
	e.mux.RLock()
	defer e.mux.RUnlock()
	return e.configs[key]
}

// SetCacheAdapter 设置缓存
func (e *Application) SetCacheAdapter(c cache.AdapterCache) {
	e.cache = c
}

// GetCacheAdapter 获取缓存
func (e *Application) GetCacheAdapter() cache.AdapterCache {
	return e.cache
}

// SetLockerAdapter 设置分布式锁
func (e *Application) SetLockerAdapter(c locker.AdapterLocker) {
	e.locker = c
}

// GetLockerAdapter 获取分布式锁
func (e *Application) GetLockerAdapter() locker.AdapterLocker {
	return e.locker
}
