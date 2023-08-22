package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/v1lezz/todo/pkg/service"
)

type emailSender interface {
	SendMessage(string, string) error
}

type Handler struct {
	services    *service.Service
	emailClient *emailSender
}

func NewHandler(services *service.Service, emailClient emailSender) *Handler {
	return &Handler{
		services:    services,
		emailClient: &emailClient,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.singIn)
		auth.POST("/sign-up", h.singUp)
	}
	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/", h.getAllTasks)
			tasks.GET("/:id", h.getTaskById)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
	}
	return router
}
