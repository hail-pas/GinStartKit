package constant

type BreakError struct {
}

func (receiver BreakError) Error() string {
	return ErrorBreakRequest
}

const (
	ErrorBreakRequest = "break request" // 中断请求
)
