package meta

// Response 统一 HTTP JSON 响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(data interface{}) Response {
	return Response{Code: 0, Message: "ok", Data: data}
}

func Fail(code int, msg string) Response {
	return Response{Code: code, Message: msg}
}
