/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2022/11/24 20:52
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2022/11/24 20:52
 */

package pbtrans

import (
	"encoding/json"
	pbstruct "github.com/golang/protobuf/ptypes/struct"
)

func PbStructToInterface(source *pbstruct.Struct, result interface{}) error {
	bytes, err := source.MarshalJSON()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, &result); err != nil {
		return err
	}
	return nil
}

func InterfaceToPbStruct(source interface{}, result **pbstruct.Struct) error {
	marshal, err := json.Marshal(source)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(marshal, &result); err != nil {
		return err
	}
	return nil
}
