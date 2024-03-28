/*
 * @Author: yujiajie
 * @Date: 2024-03-13 10:53:38
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 17:56:50
 * @FilePath: /stage/sdk/server/app.go
 * @Description:
 */
package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"stage/sdk/core"
	"stage/sdk/database"
	"stage/sdk/storage"
	"syscall"
	"time"

	"github.com/bird-coder/manyo/config"
	"github.com/bird-coder/manyo/pkg/rungroup"
)

func Init() error {
	var g rungroup.Group
	{
		term := make(chan os.Signal, 1)
		signal.Notify(term, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		ticker := time.NewTicker(time.Second * 30)
		closeChan := make(chan struct{})

		g.Add(
			func() error {
				for {
					select {
					case s := <-term:
						fmt.Println("get a signal:", s.String())
						return nil
					case <-closeChan:
						return nil
					case <-ticker.C:
						fmt.Println("running...")
					}
				}
			},
			func(err error) {
				close(closeChan)
			},
		)
	}
	{
		dbCancel := make(chan struct{})
		g.Add(
			func() error {
				if err := database.Setup(); err != nil {
					return err
				}
				<-dbCancel

				return nil
			},
			func(err error) {
				close(dbCancel)
			},
		)
	}
	{
		storageCancel := make(chan struct{})
		g.Add(
			func() error {
				if err := storage.Setup(); err != nil {
					return err
				}
				<-storageCancel

				return nil
			},
			func(err error) {
				close(storageCancel)
			},
		)
	}
	{
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cfg := core.App.GetConfig("http").(*config.HttpConfig)
		httpServer := NewHttpServer(cfg)
		httpServer.WithContext(ctx)
		g.Add(
			func() error {
				httpServer.Init()
				err := httpServer.Run()
				return err
			},
			func(err error) {
				httpServer.Close()
				cancel()
			},
		)
	}
	if err := g.Run(); err != nil {
		fmt.Println(err)
	}
	return nil
}
