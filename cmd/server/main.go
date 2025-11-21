package main

import (
	"fmt"
	"go-crud/internal/config"
	"go-crud/internal/database"
	"go-crud/internal/router"
)

func main() {
	config.Load()
	database.Connect()

	r := router.RegisterRouter()

	fmt.Println("server running on port:", config.Cfg.Port)
	r.Run(":" + config.Cfg.Port)
}
