/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-10-20 11:26:31
 * @LastEditors: John Holl
 * @LastEditTime: 2021-10-20 11:26:32
 */

package response

//const (
//	OkStatus  = "00000"
//	OkMessage = "请求成功"
//
//	InvalidParam        = "40000"
//	InvalidParamMessage = "参数不合法"
//
//	UnAuth        = 403
//	UnAuthMessage = "用户认证错误"
//)

var RESPONSE_STATUS = map[string]map[string]string{
	"SUCCESS": {
		"code":    "00000",
		"message": "请求成功",
	},
	"FAILURE": {
		"code":    "99999",
		"message": "请求失败",
	},
	"NOT_LOGGED_IN": {
		"code":    "23000",
		"message": "未登录",
	},
}

const (
	// SUCCESS_CODE 请求成功
	SUCCESS_CODE    = "00000"
	SUCCESS_MESSAGE = "请求成功"

	FAILURE_CODE    = "99999"
	FAILURE_MESSAGE = "请求失败"

	VERIFICATION_CODE_SEND_FAILED_CODE    = "10000"
	VERIFICATION_CODE_SEND_FAILED_MESSAGE = "验证码发送失败"
	/**
	 * 验证码错误
	 */
	VERIFICATION_CODE_ERROR = "10100"
	/**
	 * 验证码失效
	 */
	VERIFICATION_CODE_INVALID = "10200"
	/**
	 * 请求类型错误
	 */
	REQUEST_TYPE_ERROR = "11100"
	/**
	 * 不支持该请求类型
	 */
	REQUEST_TYPE_DENIED = "11200"
	/**
	 * 文件上传失败
	 */
	FILE_UPLOAD_FAILED = "12000"
	/**
	 * 注册登录错误
	 */
	REGISTER_OR_LOGIN_ERROR = "20000"
	/**
	 * 账号或密码不能为空
	 */
	USERNAME_OR_PASSWORD_NOT_BE_EMPTY = "20100"
	/**
	 * 认证code不能为空
	 */
	AUTH_CODE_NOT_BE_EMPTY = "20200"
	/**
	 * 认证code已失效
	 */
	AUTH_CODE_INVALID = "20300"
	/**
	 * 注册失败
	 */
	REGISTER_FAILED = "21000"
	/**
	 * 账号已存在
	 */
	ACCOUNT_ALREADY_EXISTS = "21100"
	/**
	 * 登陆失败
	 */
	LOGIN_FAILED = "22000"
	/**
	 * 账号或密码错误
	 */
	USERNAME_OR_PASSWORD_ERROR = "22100"
	/**
	 * 账号已被锁定
	 */
	ACCOUNT_IS_LOCKED = "22200"
	/**
	 * 账号未激活
	 */
	ACCOUNT_NOT_ACTIVATED = "22300"
	/**
	 * 需要绑定手机号
	 */
	NEED_BINDING_PHONE_NUMBER = "22400"
	/**
	 * 账号不存在
	 */
	ACCOUNT_NOT_EXISTS = "22500"
	/**
	 * 未登录
	 */
	NOT_LOGGED_IN_CODE    = "23000"
	NOT_LOGGED_IN_MESSAGE = "未登录"
	/**
	 * 无token
	 */
	TOKEN_IS_MISSING = "23100"
	/**
	 * token无效
	 */
	TOKEN_INVALID_CODE    = 403
	TOKEN_INVALID_MESSAGE = "token无效"
	/**
	 * token已过期
	 */
	TOKEN_EXPIRED_CODE    = "23300"
	TOKEN_EXPIRED_MESSAGE = "token已过期"
	/**
	 * token被踢下线
	 */
	TOKEN_REMOVED = "23400"
	/**
	 * token被顶下线
	 */
	TOKEN_REPLACED = "23500"
	/**
	* License已过期
	 */
	LICENSE_EXPIRED_CODE    = "24400"
	LICENSE_EXPIRED_MESSAGE = "LICENSE已过期"
	/**
	 * 数据错误
	 */
	DATA_ERROR = "30000"
	/**
	 * 数据已存在
	 */
	DATA_ALREADY_EXISTS = "30100"
	/**
	 * 数据不存在
	 */
	DATA_NOT_EXISTS = "30200"
	/**
	 * 新增错误
	 */
	ADD_FAILED = "31000"
	/**
	 * 删除错误
	 */
	DELETE_FAILED = "32000"
	/**
	 * 修改错误
	 */
	UPDATE_FAILED = "33000"
	/**
	 * 查询错误
	 */
	QUERY_FAILED = "34000"
	/**
	 * 格式错误
	 */
	PARAM_FORMAT_ERROR = "35000"
	/**
	 * 类型错误
	 */
	PARAM_TYPE_ERROR = "36000"
	/**
	 * 参数不能为空
	 */
	PARAM_NOT_BE_EMPTY_CODE    = "37000"
	PARAM_NOT_BE_EMPTY_MESSAGE = "参数不能为空"
	/**
	 * 参数范围错误
	 */
	PARAM_OUT_OF_RANGE_CODE    = "38000"
	PARAM_OUT_OF_RANGE_MESSAGE = "参数范围错误"
	/**
	 * 参数校验错误
	 */
	PARAM_CHECK_EXCEPTION_CODE    = "39000"
	PARAM_CHECK_EXCEPTION_MESSAGE = "参数校验错误"
	/**
	 * 没有权限
	 */
	PERMISSION_DENIED = "40000"
	/**
	 * 服务器错误
	 */
	EXCEPTION_CODE    = "50000"
	EXCEPTION_MESSAGE = "服务器错误"
	/**
	 * 临时停服
	 */
	TEMPORARILY_STOP_SERVICE = "50100"
	/**
	 * 服务升级中
	 */
	SERVICE_UPDATING = "50200"
)
