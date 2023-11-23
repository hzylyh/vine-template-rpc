/**
 * @Description: 告警服务
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 15:42
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 15:42
 */

package alarm

import (
	"github.com/google/wire"
	"vine-template-rpc/internal/alarm/biz"
	"vine-template-rpc/internal/alarm/data"
	"vine-template-rpc/internal/alarm/service"
)

// AlarmProviderSet is service providers.
var AlarmProviderSet = wire.NewSet(
	// service
	service.NewAlarmService,

	// biz
	biz.NewEnginBiz,

	// data common
	data.NewData,
	data.NewEntClient,

	//data biz
	//data.NewSiteRepo,
)
