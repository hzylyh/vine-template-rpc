/**
 * @Description: 鉴权
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/26 18:44
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/26 18:44
 */

package authz

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	stringadapter "github.com/casbin/casbin/v2/persist/string-adapter"
	entadapter "github.com/casbin/ent-adapter"
	"github.com/google/wire"
	"vine-template-rpc/internal/conf"
)

var AuthProviderSet = wire.NewSet(
	NewAuthz,
)

func NewAuthz(c *conf.Data) *casbin.Enforcer {
	modelText := `
[request_definition]
r = sub, sub2, act, obj

[policy_definition]
p = sub, sub2, act, obj

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (r.sub == p.sub || p.sub == "*") && (r.sub2 == p.sub2 || p.sub2 == "*") && r.act == p.act && r.obj == p.obj
`
	a, err := entadapter.NewAdapter(c.Database.Driver, fmt.Sprintf("%s", c.Database.Source)) // Your driver and data source.
	if err != nil {
		panic(err)
	}

	m, err := model.NewModelFromString(modelText)
	if err != nil {
		panic(err)
	}

	Enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}

	Enforcer.ClearPolicy()

	ruleText := `
p, *, auth, read, /api.system.v1.System/AddUser
p, *, auth, read, /api.system.v1.System/UpdateUser
p, *, auth, read, /api.system.v1.System/DeleteUser
p, *, auth, read, /api.system.v1.System/GetUser
p, *, auth, read, /api.system.v1.System/ListUser

p, *, auth, read, /api.system.v1.System/AddRole
p, *, auth, read, /api.system.v1.System/BindUser

p, *, auth, read, /api.system.v1.System/AddPerm

p, *, *, read, /api.system.v1.System/Login

p, *, auth, read, /api.emonitor.v1.Emonitor/AddSite
p, *, auth, read, /api.alarm.v1.Alarm/Trigger
`

	sa := stringadapter.NewAdapter(ruleText)
	// load all rules from string adapter to enforcer's memory
	err = sa.LoadPolicy(Enforcer.GetModel())
	if err != nil {
		panic(err)
	}

	// save all rules from enforcer's memory to Xorm adapter (DB)
	// same as:
	// a.SavePolicy(Enforcer.GetModel())
	err = Enforcer.SavePolicy()
	if err != nil {
		panic(err)
	}

	return Enforcer
}
