/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 15:50
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 15:50
 */

package biz

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type RuleBiz struct {
	log *log.Helper
}

func (rb *RuleBiz) Add(info string) {
	fmt.Println(info)
}

func NewRuleBiz(
	logger log.Logger,
) *RuleBiz {
	return &RuleBiz{
		log: log.NewHelper(log.With(logger, "module", "biz/rule")),
	}
}
