/*
 * @Author: yujiajie
 * @Date: 2024-03-13 14:46:09
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 17:55:09
 * @FilePath: /stage/sdk/core/type.go
 * @Description:
 */
package core

import (
	"stage/sdk/storage/cache"
	"stage/sdk/storage/locker"

	"gorm.io/gorm"
)

type Core interface {
	SetDb(key string, db *gorm.DB)
	GetDb(key string) *gorm.DB
	GetAllDb() map[string]*gorm.DB

	// SetLogger(log logger.Logger)
	// GetLogger() logger.Logger

	SetConfig(key string, config any)
	GetConfig(key string) any

	SetCacheAdapter(cache.AdapterCache)
	GetCacheAdapter() cache.AdapterCache

	SetLockerAdapter(locker.AdapterLocker)
	GetLockerAdapter() locker.AdapterLocker
}
