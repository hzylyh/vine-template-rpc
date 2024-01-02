/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 10:04
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 10:04
 */

package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Data struct {
	gdb *gorm.DB
}

// NewData .
func NewData(
	gdb *gorm.DB,
	logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "order-service/data"))

	d := &Data{
		//db:  entClient,
		gdb: gdb,
		//redisCli: redisCmd,
	}
	return d, func() {
		log.Info("closing the data resources")
	}, nil
}
