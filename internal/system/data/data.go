/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/16 14:06
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/16 14:06
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"vine-template-rpc/internal/conf"
	"vine-template-rpc/internal/emonitor/data/schema"
)

//// ProviderSet is data providers.
//var ProviderSet = wire.NewSet(
//	NewData,
//	NewEntClient,
//	NewUserRepo,
//)

// Data .
type Data struct {
	//db *ent.Client
	gdb *gorm.DB
	rdb *redis.Client
}

// NewEntClient 创建ent客户端
//func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
//	log := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
//
//	client, err := ent.Open(
//		conf.Database.Driver,
//		conf.Database.Source,
//	)
//	if err != nil {
//		log.Fatalf("failed opening connection to db: %v", err)
//	}
//	// Run the auto migration tool.
//	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
//		log.Fatalf("failed creating schema resources: %v", err)
//	}
//	return client
//}

func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {
	log := log.NewHelper(log.With(logger, "module", "system-service/data/redis"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
	})
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalf("failed opening connection to redis: %v", err)
	}
	return rdb
}

// NewGormDB gorm数据库
func NewGormDB(dialector *gorm.Dialector, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "system-service/data/gorm"))

	db, err := gorm.Open(*dialector, &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Info),
	})
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	if err := db.AutoMigrate(
		schema.Schemas...,
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return db
}

// NewMysqlDialector Mysql Dialector
func NewMysqlDialector(conf *conf.Data) *gorm.Dialector {
	open := mysql.Open(conf.Database.Source)
	return &open
}

// NewData .
func NewData(
	gdb *gorm.DB,
	rdb *redis.Client,
	logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "system-service/data"))

	d := &Data{
		gdb: gdb,
		rdb: rdb,
	}
	return d, func() {
		log.Info("closing the data resources")
	}, nil
}
