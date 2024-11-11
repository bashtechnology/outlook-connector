package main

import (
	"log"
	"time"

	"outlook-connector/api/routes"
	"outlook-connector/config"
)

// License: MIT http://opensource.org/licenses/MIT//
// @title outlook-connector
// @version 1.0
// @description Api integração connector de emails outlook
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http, https
func main() {
	loadConfig, err := config.LoadConfig("env")
	if err != nil {
		log.Fatal("🚀 Falha ao carregar modulos Env: ", err)
	}
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	router := routes.Init(loadConfig)
	err = router.Run(":" + loadConfig.ServerPort)
	if err != nil {
		log.Fatal("🚀 Falha ao iniciar o servidor: ", err)
	}
}
