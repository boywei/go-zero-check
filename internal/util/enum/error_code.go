package enum

// 定义错误码
type ErrCode struct {
	Code    int
	Message string
}

func (err ErrCode) Error() string {
	return err.Message
}
