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
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"vine-template-rpc/internal/conf"
	"vine-template-rpc/internal/emonitor/data/ent"
	"vine-template-rpc/internal/emonitor/data/ent/migrate"
)

// Data .
type Data struct {
	db *ent.Client
	//redisCli redis.Cmdable
}

// NewEntClient 创建ent客户端
func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "emonitor-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(
	entClient *ent.Client,
	//redisCmd redis.Cmdable,
	logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "emonitor-service/data"))

	d := &Data{
		db: entClient,
		//redisCli: redisCmd,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
