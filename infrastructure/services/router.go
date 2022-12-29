package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler func() (any, error)

type Router interface {
	Get(url string, handler Handler)
	Start(port int) error
}

func NewRouter() Router {
	engine := gin.Default()
	_ = engine.SetTrustedProxies(nil)
	return &RouterGin{engine: engine}
}

type RouterGin struct {
	engine *gin.Engine
}

func (r *RouterGin) Get(url string, handler Handler) {
	r.engine.GET(url, func(c *gin.Context) {
		result, err := handler()
		if err != nil {
			switch v := err.(type) {
			case HttpError:
				_ = c.AbortWithError(v.code, v)
			default:
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}
			return
		}
		c.JSON(http.StatusOK, result)
	})
}

func (r *RouterGin) Start(port int) error {
	return r.engine.Run("localhost:" + strconv.Itoa(port))
}

func NullRouter() *RouterNull {
	return &RouterNull{getHandlers: make(map[string]Handler)}
}

type RouterNull struct {
	getHandlers map[string]Handler
}

func (r *RouterNull) Start(port int) error {
	return nil
}

func (r *RouterNull) Get(url string, handler Handler) {
	r.getHandlers[url] = handler
}

func (r *RouterNull) Call(method string, url string) (any, error) {
	return r.getHandlers[url]()
}
