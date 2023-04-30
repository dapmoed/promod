package pm_servers

import (
	"github.com/gin-gonic/gin"
	"promod/pkg/app"
)

type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}

type Handler interface {
	RegisterRoutes(router *gin.Engine)
	SetLogger(logger app.Logger) Handler
}

type GinServer struct {
	Router   *gin.Engine
	Handlers []Handler
}

func NewGinServer(handlers ...Handler) *GinServer {
	return &GinServer{
		Router:   gin.Default(),
		Handlers: handlers,
	}
}

func (g *GinServer) Serve(logger app.Logger) error {
	for _, h := range g.Handlers {
		h.SetLogger(logger).RegisterRoutes(g.Router)
	}
	return g.Router.Run()
}
