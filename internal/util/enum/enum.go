package enum

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

var (
	Success = &ErrCode{Code: 200, Message: "SUCCESS"}
	Failure = &ErrCode{Code: 400, Message: "FAILURE"}

	// 系统错误, 前缀为 100
	InternalServerError = &ErrCode{Code: 10001, Message: "内部服务器错误"}
	InvalidParam        = &ErrCode{Code: 10002, Message: "请求参数错误"}
	ErrSyntax           = &ErrCode{Code: 10003, Message: "模型语法错误"}

	// 模型错误, 前缀是 202
	ModelConvertErr = &ErrCode{Code: 20201, Message: "模型解析错误"}
	ModelNotExist   = &ErrCode{Code: 20202, Message: "模型不存在"}
	ModelRunErr     = &ErrCode{Code: 20203, Message: "模型运行失败"}
)
