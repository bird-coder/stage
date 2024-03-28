/*
 * @Author: yujiajie
 * @Date: 2024-02-23 17:11:37
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 18:07:16
 * @FilePath: /stage/sdk/server/http.go
 * @Description:
 */
package server

import (
	"context"
	"fmt"
	"net/http"
	"stage/app/routes"
	"stage/internal/middleware"
	"time"

	"github.com/bird-coder/manyo/config"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*http.Server

	cfg *config.HttpConfig
	ctx context.Context
}

func NewHttpServer(cfg *config.HttpConfig) *HttpServer {
	srv := &HttpServer{cfg: cfg}
	return srv
}

func (srv *HttpServer) Init() {
	r := gin.New()
	middleware.Init(r)
	routes.InitRouter(r)

	srv.Server = &http.Server{
		Addr:           srv.cfg.Addr,
		ReadTimeout:    time.Duration(srv.cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(srv.cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: srv.cfg.MaxHeaderBytes,
		Handler:        r,
	}
}

func (srv *HttpServer) WithContext(ctx context.Context) {
	srv.ctx = ctx
}

func (srv *HttpServer) Run() error {
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("http server exit:", err)
		if err == http.ErrServerClosed {
			fmt.Println("waiting for http server shutdown finish...")
		} else {
			fmt.Println("http server出错, err: ", err)
		}
		return err
	}
	return nil
}

func (srv *HttpServer) Close() error {
	fmt.Println("http server shutdown begin。。。")
	if err := srv.Shutdown(srv.ctx); err != nil {
		fmt.Println("http server shutdown error:", err)
		return err
	}
	fmt.Println("http server shutdown processed successfully")
	return nil
}
