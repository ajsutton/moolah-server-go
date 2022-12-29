package services

type Handler func() (any, error)

type Router interface {
	Get(url string, handler Handler)
}

func NewRouter() Router {
	return NullRouter()
}

func NullRouter() *RouterNull {
	return &RouterNull{getHandlers: make(map[string]Handler)}
}

type RouterNull struct {
	getHandlers map[string]Handler
}

func (r *RouterNull) Get(url string, handler Handler) {
	r.getHandlers[url] = handler
}

func (r *RouterNull) Call(method string, url string) (any, error) {
	return r.getHandlers[url]()
}
