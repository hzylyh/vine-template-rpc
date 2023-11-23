/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/11/21 10:48
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/11/21 10:48
 */

package judge

import (
	"encoding/json"
	"fmt"
	"strconv"
	"vine-template-rpc/pkg/jsonpath"
)

type Rule struct {
	Path string `json:"path"`
	Op   string `json:"op"` // > < = != in notin
	Val  string `json:"val"`
}

type Judge struct {
	// 规则
	Rule Rule
	// 数据
	Data string
}

func NewJudge(rule Rule, data string) *Judge {
	return &Judge{
		Rule: rule,
		Data: data,
	}
}

func (j *Judge) Judge() (bool, error) {
	var jsonData interface{}
	err := json.Unmarshal([]byte(j.Data), &jsonData)
	if err != nil {
		return false, err
	}
	res, err := jsonpath.JsonPathLookup(jsonData, j.Rule.Path)
	if err != nil {
		return false, err
	}
	switch res.(type) {
	case int:
		exprVal, err := strconv.Atoi(j.Rule.Val)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return compareInt(res.(int), exprVal), nil
	case string:
		return compareString(res.(string), j.Rule.Val), nil
	}
	return false, err
}

func compareInt(act, exp int) bool {
	return false
}

func compareString(act, exp string) bool {
	return false
}
