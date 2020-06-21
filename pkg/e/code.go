package e

const (
	// 响应成功
	SUCCESS = 200
	// 响应错误
	ERROR = 500
	// 错误请求参数
	INVALID_PARAMS = 400

	// TAG已存在
	ERROR_EXIST_TAG = 10001
	// TAG不存在
	ERROR_NOT_EXIST_TAG = 10002
	// 文章不存在
	ERROR_NOT_EXIST_ARTICLE = 10003

	// TOKEN无效
	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
	// TOKEN超时
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	// TOKEN错误
	ERROR_AUTH_TOKEN = 20003
	// 无效用户
	ERROR_AUTH = 20004
	// 用户已存在
	ERROR_EXIST_AUTH = 20005
	// 密码错误
	ERROR_AUTH_PASSWORD = 20006

	// 上传图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL = 30001
	// 图片检查失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL = 30002
	// 图片校验错误
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)
