package main

import (
	"github.com/achanda/gorest/app"
	"github.com/achanda/gorest/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
