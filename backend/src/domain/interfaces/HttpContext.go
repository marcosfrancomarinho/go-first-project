package interfaces

type HttpContext interface {
	GetBody(input any) error
	Send(status int, datas any, token ...string)
	SendError(err error)
	GetToken() string
	SetIdentifiers(key string, datas any)
	GetIdentifiers(key string) (any, error)
	GetParams(key string) (*string, error)
}
