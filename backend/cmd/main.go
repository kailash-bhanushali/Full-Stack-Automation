package main

import (
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/internal/config"
	"github.com/kailash-bhanushali/Full-Stack-Automation/backend/pkg/server"
)

func main() {
	server.NewServer(config.NewServerConfig())
}
