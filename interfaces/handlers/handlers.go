package handlers

import (
	"fmt"
	"lrucache/application/services"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Run(cfg HandlersCfg)
}

type HandlersCfg struct {
	ServePort int `arg:"--serve-port,env:SERVE_PORT" default:"8080"`
}

type handlersImpl struct {
	router *gin.Engine

	userHandlers UserHandlers

	services services.Services
}

func NewHandlers(services services.Services) Handlers {
	h := &handlersImpl{
		router:       gin.Default(),
		userHandlers: NewUserHandlers(services),
	}

	h.router.POST("/user", h.userHandlers.CreateUser)
	h.router.GET("/user/:email", h.userHandlers.GetUser)

	return h
}

var _ Handlers = (*handlersImpl)(nil)

func (h *handlersImpl) Run(cfg HandlersCfg) {
	_ = h.router.Run(fmt.Sprintf(":%d", cfg.ServePort))
}
