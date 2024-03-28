/*
 * @Author: yujiajie
 * @Date: 2024-03-15 15:40:40
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 15:40:52
 * @FilePath: /stage/sdk/storage/cache/type.go
 * @Description:
 */
package cache

import "time"

type AdapterCache interface {
	String() string
	Get(key string) (string, error)
	Set(key string, val interface{}, expire int) error
	Del(key string) error
	HGet(hk, key string) (string, error)
	HDel(hk, key string) error
	Increase(key string) error
	Decrease(key string) error
	Expire(key string, dur time.Duration) error
}
