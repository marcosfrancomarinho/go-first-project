package interfaces

type HttpContext interface {
	GetBody(input any) ( error)
	Send(status int, datas any)
	SendError(error error)
}
