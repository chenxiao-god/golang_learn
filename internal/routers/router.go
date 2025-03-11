package routers

import (
	"github.com/gin-gonic/gin"
	"gomod/internal/handle"
	"gomod/internal/middleware"
)

func NewRouter(userHandler *handle.UserHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.StatCost())
	public := router.Group("/api/user")
	{
		public.POST("/create", userHandler.Creat)
	}
	return router
}
