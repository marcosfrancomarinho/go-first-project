package interfaces

type Method string


type HttpServer interface {
	On(method string, path string, controllers HttpControllers)
	Run(port int)
}

