/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/20 16:14
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/20 16:14
 */

package biz

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"vine-template-rpc/internal/alarm/pkg/judge"
)

type EngineBiz struct {
	log *log.Helper
}

// Trigger 根据jsonpath获取数据并判断是否触发报警
func (eb *EngineBiz) Trigger(id, data string) (err error) {
	// 根据id获取topic

	// 订阅topic，获取数据
	rule := judge.Rule{
		Path: "$.store.book[0].price",
		Op:   ">",
		Val:  "10",
	}
	data = `{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": 10
}
`

	//// 根据jsonpath获取data中对应的值
	//var jsonData interface{}
	//err = json.Unmarshal([]byte(data), &jsonData)
	//if err != nil {
	//	return err
	//}
	//res, err := jsonpath.JsonPathLookup(jsonData, rule.Path)
	//if err != nil {
	//	return err
	//}
	//// 根据id获取rule
	//fmt.Println(res)

	// 根据rule中的规则判断是否触发报警
	cm := judge.NewJudge(rule, data)
	jres, _ := cm.Judge()
	if jres {
		triggerAlarm()
	}

	return
}

func triggerAlarm() {
	fmt.Println("dddd")
}

func NewEnginBiz(logger log.Logger) *EngineBiz {
	return &EngineBiz{
		log: log.NewHelper(log.With(logger, "module", "biz/engine")),
	}
}
