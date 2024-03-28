/*
 * @Author: yujiajie
 * @Date: 2024-03-15 15:41:05
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 15:41:08
 * @FilePath: /stage/sdk/storage/locker/type.go
 * @Description:
 */
package locker

import "github.com/bsm/redislock"

type AdapterLocker interface {
	String() string
	Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error)
}
