package main

import (
	"fmt"
	"gomod/internal/config"
	"gomod/internal/database"
	"gomod/internal/handle"
	"gomod/internal/repository"
	"gomod/internal/routers"
	"gomod/internal/service"
)

func main() {

	cfg, _ := config.Load("application-dev.yaml")
	db, _ := database.NewDB(&cfg.DB)
	// 初始化各层
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(&userRepo)
	userHandler := handle.NewUserHandler(userService)
	router := routers.NewRouter(userHandler)
	// 启动服务
	router.Run(fmt.Sprintf(":%d", cfg.HTTP.Port))
}
