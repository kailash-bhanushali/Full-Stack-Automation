package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/internal/config"
	"net/http"
)

type statusSVC struct{
	router *gin.Engine
	Env string
}

func statusSVCHandler(mainHandler Handler, groupBase *gin.RouterGroup){
	handler := statusSVC{
		router: mainHandler.router,
		Env: mainHandler.env,
	}
	group := groupBase.Group("")
	group.GET("/ping", handler.ping)
	group.GET("/config", handler.config)
}

func (h *statusSVC) ping(c *gin.Context){
	c.JSON(http.StatusOK, "Service Up")
}

func (h *statusSVC) config(c *gin.Context){
	c.JSON(http.StatusOK, fmt.Sprintf("Env Config Loaded: %#v", config.NewServerConfig()))
}