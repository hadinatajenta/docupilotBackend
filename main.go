package main

import (
	"fmt"
	"gobackend/app"
	"gobackend/core/configuration"
	"gobackend/core/contract"
	"gobackend/core/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configuration.MustLoad()

	deps := app.InitDependencies(cfg)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	api := r.Group("api")
	contract.RegisterRoutes(api, deps)

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting server at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
