/*
 * @Author: yujiajie
 * @Date: 2024-02-23 17:11:50
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-13 09:25:19
 * @FilePath: /stage/server/server.go
 * @Description:
 */
package server

type Server interface {
	init()
	Run() error
	Close() error
}
