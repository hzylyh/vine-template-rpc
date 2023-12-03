/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 10:39
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 10:39
 */

package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vine-template-rpc/internal/conf"
	"vine-template-rpc/internal/emonitor/data/schema"
)

// Data .
type Data struct {
	//db  *ent.Client
	gdb *gorm.DB
	//redisCli redis.Cmdable
}

// NewEntClient 创建ent客户端
//func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
//	log := log.NewHelper(log.With(logger, "module", "emonitor-service/data/ent"))
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

// NewGormDB gorm数据库
func NewGormDB(dialector *gorm.Dialector, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "emonitor-service/data/gorm"))

	db, err := gorm.Open(*dialector, &gorm.Config{
		//Logger: logger.,
	})
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	if err := db.AutoMigrate(schema.Schemas...); err != nil {
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
	//entClient *ent.Client,
	gdb *gorm.DB,
	//redisCmd redis.Cmdable,
	logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "emonitor-service/data"))

	d := &Data{
		//db:  entClient,
		gdb: gdb,
		//redisCli: redisCmd,
	}
	return d, func() {
		log.Info("closing the data resources")
	}, nil
}
