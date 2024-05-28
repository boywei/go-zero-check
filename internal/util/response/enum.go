package response

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

var (
	SuccessResp = &ErrCode{Code: 200, Message: "SUCCESS"}
	FailureResp = &ErrCode{Code: 400, Message: "FAILURE"}

	// 系统错误, 前缀是100
	InternalServerError = &ErrCode{Code: 10001, Message: "内部服务器错误"}
	InvalidParam        = &ErrCode{Code: 10002, Message: "请求参数错误"}
	ErrSyntax           = &ErrCode{Code: 10003, Message: "模型语法错误"}

	// 模型错误, 前缀是202
	ModelConvertErr = &ErrCode{Code: 20201, Message: "模型解析错误"}
	ModelNotExist   = &ErrCode{Code: 20202, Message: "模型不存在"}
	ModelRunErr     = &ErrCode{Code: 20203, Message: "模型运行失败"}

	// SynLong/Lustre错误, 前缀是201
	SynLongConvertErr = &ErrCode{Code: 20101, Message: "SynLong模型转换失败"}
	SynLongDataflowErr = &ErrCode{Code: 20102, Message: "SynLong数据流验证失败"}

	// 验证器错误, 前缀是203
	ModelVerifyErr = &ErrCode{Code: 20301, Message: "模型验证错误"}
)
