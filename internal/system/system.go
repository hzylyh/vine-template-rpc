/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/16 14:41
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/16 14:41
 */

package system

import (
	"github.com/google/wire"
	"vine-template-rpc/internal/system/biz"
	"vine-template-rpc/internal/system/data"
	"vine-template-rpc/internal/system/service"
)

// SystemProviderSet is service providers.
var SystemProviderSet = wire.NewSet(
	service.NewSystemService,
	biz.NewUserBiz,
	data.NewData,
	data.NewEntClient,
	data.NewUserRepo,
)
