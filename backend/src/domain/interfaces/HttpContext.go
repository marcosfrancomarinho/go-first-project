package interfaces

type HttpContext interface {
	GetBody(input any) error
	Send(status int, datas any, token ...string)
	SendError(err error)
	GetToken() string
	GetParams(key string) (*string, error)
	GetQuery(key ...string) (map[string]string, error)
}
