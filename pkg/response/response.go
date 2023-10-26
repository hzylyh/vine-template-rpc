/**
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2022/11/26 14:51
 * @LastEditors: John Holl
 * @LastEditTime: 2022/11/26 14:51
 */

package response

import (
	"fmt"
)

type Response struct {
	Code   string      `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

type HTTPError struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

func (r *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError code: %s message: %s", r.Code, r.Msg)
}
