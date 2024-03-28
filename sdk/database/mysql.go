/*
 * @Author: yujiajie
 * @Date: 2024-03-14 15:53:39
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 17:23:58
 * @FilePath: /stage/sdk/database/mysql.go
 * @Description:
 */
package database

import (
	"fmt"
	"stage/sdk/core"

	"github.com/bird-coder/manyo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func Setup() error {
	dbConfigs := core.App.GetConfig("databases").(map[string]*config.MysqlConfig)
	for k, cfg := range dbConfigs {
		if err := setupDatabase(k, cfg); err != nil {
			return err
		}
	}
	return nil
}

func setupDatabase(host string, c *config.MysqlConfig) error {
	fmt.Println("mysql connection: ", c.Default)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.Default,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Println("mysql init failed")
		return err
	}
	if c.Cluster {
		initCluster(db, c)
	}
	core.App.SetDb(host, db)
	return nil
}

func initCluster(db *gorm.DB, c *config.MysqlConfig) {
	if len(c.Sources) == 0 && len(c.Replicas) == 0 {
		return
	}

	config := dbresolver.Config{}
	config.Sources = []gorm.Dialector{}
	config.Replicas = []gorm.Dialector{}

	for _, sourceConfig := range c.Sources {
		config.Sources = append(config.Sources, mysql.Open(sourceConfig))
	}
	for _, replicaConfig := range c.Replicas {
		config.Replicas = append(config.Replicas, mysql.Open(replicaConfig))
	}
	db.Use(dbresolver.Register(config))
}
