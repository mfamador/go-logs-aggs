package main

import (
	"github.com/mfamador/go-logs-aggs/app"
	"github.com/mfamador/go-logs-aggs/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
