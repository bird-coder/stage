/*
 * @Author: yujiajie
 * @Date: 2024-03-15 15:06:04
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 17:25:02
 * @FilePath: /stage/sdk/storage/initialize.go
 * @Description:
 */
package storage

import (
	"fmt"
	"stage/sdk/core"
	"stage/sdk/storage/cache"
	"stage/sdk/storage/locker"
	"time"

	"github.com/bird-coder/manyo/config"
	"github.com/redis/go-redis/v9"
)

var (
	_redis *redis.Client
)

func Setup() error {
	cacheConfig := core.App.GetConfig("cache").(*config.CacheConfig)
	if cacheConfig != nil {
		if err := setupCache(cacheConfig); err != nil {
			return err
		}
	}
	lockConfig := core.App.GetConfig("locker").(*config.LockConfig)
	if lockConfig != nil {
		if err := setupLocker(lockConfig); err != nil {
			return err
		}
	}
	return nil
}

func setupCache(cfg *config.CacheConfig) error {
	if cfg.Redis != nil {
		options := redis.Options{
			Addr:         cfg.Redis.Addr,
			Password:     cfg.Redis.Password,
			DB:           cfg.Redis.Db,
			Protocol:     cfg.Redis.Protocol,
			DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
			PoolSize:     cfg.Redis.PoolSize,
			MinIdleConns: cfg.Redis.IdleConns,
		}
		r, err := cache.NewRedis(nil, &options)
		_redis = r.GetClient()
		if err != nil {
			fmt.Println("cache setup error", err)
			return err
		}
		core.App.SetCacheAdapter(r)
	}
	return nil
}

func setupLocker(cfg *config.LockConfig) error {
	if cfg.Redis != nil {
		var options *redis.Options
		if _redis == nil {
			options = &redis.Options{
				Addr:         cfg.Redis.Addr,
				Password:     cfg.Redis.Password,
				DB:           cfg.Redis.Db,
				Protocol:     cfg.Redis.Protocol,
				DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
				ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
				WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
			}
		}
		r, err := locker.NewRedis(_redis, options)
		if err != nil {
			fmt.Println("locker setup error", err)
			return err
		}
		core.App.SetLockerAdapter(r)
	}
	return nil
}
