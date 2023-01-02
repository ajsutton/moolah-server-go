package services

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Handler func() (any, error)

type Router interface {
	Get(url string, handler Handler)
	Post(url string, handler Handler)
	Start(port int) error
	Stop() error
	Call(method string, url string) (int, string, error)
}

func NewRouter() Router {
	engine := gin.Default()
	_ = engine.SetTrustedProxies(nil)
	return &RouterGin{engine: engine}
}

type RouterGin struct {
	engine *gin.Engine
	server *http.Server
}

func (r *RouterGin) Call(method string, url string) (int, string, error) {
	fullUrl := "http://" + r.server.Addr + url
	var resp *http.Response
	var err error
	switch method {
	case http.MethodGet:
		resp, err = http.Get(fullUrl)
	case http.MethodPost:
		resp, err = http.Post(
			fullUrl, "application/json", bytes.NewBufferString(""))
	default:
		return 0, "", CallError("Unknown method type: " + method)
	}
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	return resp.StatusCode, string(body), nil
}

func (r *RouterGin) Get(url string, handler Handler) {
	r.engine.GET(url, handlerWrapper(handler))
}

func (r *RouterGin) Post(url string, handler Handler) {
	r.engine.POST(url, handlerWrapper(handler))
}

func handlerWrapper(handler Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := handler()
		if err != nil {
			switch v := err.(type) {
			case HttpError:
				c.AbortWithStatusJSON(v.code, v.message)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			}
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (r *RouterGin) Start(port int) error {
	r.server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r.engine,
	}
	ch := make(chan int)
	go func() {
		ln, err := net.Listen("tcp", r.server.Addr)
		if err != nil {
			log.Fatal("Failed to start listening", err)
			return
		}
		r.server.Addr = ln.Addr().String()
		ch <- 0
		err = r.server.Serve(ln)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start serving", err)
			return
		}
	}()
	<-ch
	return nil
}

func (r *RouterGin) Stop() error {
	return r.server.Close()
}

type CallError string

func (b CallError) Error() string {
	return string(b)
}
