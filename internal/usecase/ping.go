package usecase

import (
	"github.com/gin-gonic/gin"
	"promod/pkg/app"
	"promod/pkg/pm_servers"
)

type Ping struct {
	Logger app.Logger
}

func NewPing() *Ping {
	return &Ping{}
}

func (p *Ping) RegisterRoutes(r *gin.Engine) {
	r.Handle("GET", "/ping", p.Ping)
}

func (p *Ping) SetLogger(logger app.Logger) pm_servers.Handler {
	p.Logger = logger
	return p
}

func (p *Ping) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
