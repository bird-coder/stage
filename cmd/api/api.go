/*
 * @Author: yujiajie
 * @Date: 2024-02-28 16:39:51
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 17:32:17
 * @FilePath: /stage/cmd/api/api.go
 * @Description:
 */
package api

import (
	"fmt"
	"os"
	"stage/sdk/config"
	"stage/sdk/core"
	"stage/sdk/server"

	zconfig "github.com/bird-coder/manyo/config"
	"github.com/bird-coder/manyo/constant"
	zlog "github.com/bird-coder/manyo/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "start api server",
		Example:      "stage server -c config/server.yaml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/server.yaml", "Start server with provided configuration file")
}

func setup() {
	appConfig := new(config.AppConfig)
	if err := appConfig.LoadConfig(configYml); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("starting api server...")
}

func run() error {
	logConfig := core.App.GetConfig("logger").(*zconfig.LoggerConfig)
	zlog.NewLogger(logConfig, constant.Dev.String())
	defer zlog.Sync()
	zlog.Info("stage server start")

	server.Init()

	return nil
}
