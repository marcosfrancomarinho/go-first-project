package interfaces

type HttpContext interface {
	GetBody(input any) error
	Send(status int, datas any, token ...string)
	SendError(err error)
	GetToken() string
	GetParams(input any)  error
	GetQuery(input  any) error
}
