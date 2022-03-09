package server

import (
	"fmt"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/internal/config"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/pkg"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/pkg/handler/controller"
	log "github.com/sirupsen/logrus"
	"os"
)

func NewServer(conf *config.ServerConfig){
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.WithField("%+v", conf)
	factory := pkg.NewFactory(conf)
	controller.NewHandler(factory)
	address := fmt.Sprintf(":%v", conf.Port)
	if err := factory.CreateRouterEngine().Run(address); err != nil{
		log.WithError(err).Panic("Failed to Start")
	}

}
