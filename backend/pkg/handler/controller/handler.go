package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/pkg"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/pkg/api"
)

type Handler struct {
	router *gin.Engine
	env string
}

func NewHandler(factory pkg.Factory){
	handler := Handler{
		factory.CreateRouterEngine(),
		factory.EnvRelation(),
	}
	root := handler.router.Group("")
	group := root.Group(api.HANDLERBASEGROUP)

	statusSVCHandler(handler, group)
}
