package response

// ErrCode 定义错误码
type ErrCode struct {
	Code    int    `json:"code,omitempty" example:"400"`        // 错误码
	Message string `json:"message,omitempty" example:"FAILURE"` // 消息
}

func (err ErrCode) Error() string {
	return err.Message
}
