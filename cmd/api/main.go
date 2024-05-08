package main

import (
	"github.com/spf13/viper"
	"log"
	"todo-go/config"
	"todo-go/server"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err)
	}
}
