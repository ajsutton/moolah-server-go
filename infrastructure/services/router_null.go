package services

import (
	"encoding/json"
	"net/http"
)

func NullRouter() *RouterNull {
	return &RouterNull{
		getHandlers:  make(map[string]Handler),
		postHandlers: make(map[string]Handler),
	}
}

type RouterNull struct {
	getHandlers  map[string]Handler
	postHandlers map[string]Handler
}

func (r *RouterNull) Start(port int) error {
	return nil
}

func (f *RouterNull) Stop() error {
	return nil
}

func (r *RouterNull) Get(url string, handler Handler) {
	r.getHandlers[url] = handler
}

func (r *RouterNull) Post(url string, handler Handler) {
	r.postHandlers[url] = handler
}

func (r *RouterNull) Call(method string, url string) (int, string, error) {
	var handlers map[string]Handler
	switch method {
	case http.MethodGet:
		handlers = r.getHandlers
	case http.MethodPost:
		handlers = r.postHandlers
	default:
		return 0, "", CallError("Unknown method: " + method)
	}
	result, err := handlers[url]()
	if err != nil {
		switch v := err.(type) {
		case HttpError:
			return serialize(v.code, v.message)
		default:
			return serialize(http.StatusInternalServerError, err.Error())
		}
	}
	return serialize(http.StatusOK, result)
}

func serialize(status int, content any) (int, string, error) {
	serialized, err := json.Marshal(content)
	if err != nil {
		return 0, "", err
	}
	return status, string(serialized), nil
}
