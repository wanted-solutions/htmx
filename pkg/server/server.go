package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func New() *Engine {
	return &Engine{
		handler: gin.New(),
	}
}

type Engine struct {
	handler *gin.Engine
	server  *http.Server
}

func (e *Engine) Use(middleware ...gin.HandlerFunc) {
	e.handler.Use(middleware...)
}

func (e *Engine) LoadComponents(pattern ...string) {
	LoadTemplates(pattern...)
}

func (e *Engine) LoadTemplates(pattern ...string) {
	LoadTemplates(pattern...)
}

func (e *Engine) Static(relativePath, root string) {
	e.handler.Static(relativePath, root)
}

func (e *Engine) GET(relativePath string, handlers ...gin.HandlerFunc) {
	e.handler.GET(relativePath, handlers...)
}

// Public functions
func (e *Engine) ListenAndServe(addr string) error {
	templates := GetTemplates()
	e.createHTTPServer(addr)
	fmt.Println(templates)
	e.handler.LoadHTMLFiles(templates...)
	return e.server.ListenAndServe()
}

// Private functions
func (e *Engine) createHTTPServer(addr string) {
	e.server = &http.Server{
		Addr:           fmt.Sprintf(":%s", addr),
		Handler:        e.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
