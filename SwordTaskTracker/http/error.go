package http

// 以 1xxxxx 开头的错误表示 通用错误, 比如表单验证失败, 未登录等
// 以 2xxxxx 开头的错误表示 数据错误, 比如按照指定的表单查询数据库, 但数据没有找到等.
// 其他的开发过程中遇到, 再自行增加, 加完后记得改注释

// New :
// func New(code int, text string) error {
// 	return &E{code, text}
// }

// Err :
type Err struct {
	Code    int
	Message string
}

// ErrNoError : 没有错误
var ErrNoError = Err{0, "OK"}

// ErrServer : 服务端通用错误
var ErrServer = Err{10001, "Server error"}

// ErrParams : API参数出错
var ErrParams = Err{10002, "Params error"}

// ErrForm : 表单验证失败, 请查询 API 文档
var ErrForm = Err{10003, "Form error"}

// ErrNotFound : 没有该对象
var ErrNotFound = Err{10004, "Not found"}

// ErrCookie : 获取cookie失败
var ErrCookie = Err{10005, "Get cookie failed"}

// ErrFakeRequest : 伪造请求
var ErrFakeRequest = Err{10006, "Fake request"}

// ErrPermissionDenied : 没有权限
var ErrPermissionDenied = Err{10007, "Permission denied"}

// ErrLackResources : 计算资源受限
var ErrLackResources = Err{10008, "Lack of computing resources"}
