package interfaces

type Method string

type HttpServer interface {
	On(method string, path string, controllers HttpControllers, middlewares ...HttpControllers)
	Run(port int)
}
