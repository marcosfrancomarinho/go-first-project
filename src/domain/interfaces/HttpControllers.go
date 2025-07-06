package interfaces

type HttpControllers interface {
	Execute(httpContext HttpContext)
}
