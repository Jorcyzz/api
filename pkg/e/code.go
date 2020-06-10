package e

const (
	SUCCESS        = 200 //成功响应请求
	ERROR          = 500 //错误响应请求
	INVALID_PARAMS = 400 //请求参数无效

	ERROR_EXIST_TAG         = 10001 //标记错误
	ERROR_NOT_EXIST_TAG     = 10002 //错误的不存在的标记
	ERROR_NOT_EXIST_ARTICLE = 10003 //错误的不存在的文章

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 //token无效
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 //token超时
	ERROR_AUTH_TOKEN               = 20003 //taoken错误
	ERROR_AUTH                     = 20004 //无效的用户
	ERROR_EXIST_AUTH               = 20005 //手机号已存在
	ERROR_AUTH_PASSWORD            = 20006 //密码错误

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001 // 保存图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002 // 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003 // 校验图片错误，图片格式或大小有问题
)
