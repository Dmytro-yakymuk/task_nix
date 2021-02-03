package handler

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/service"
	"github.com/labstack/echo/v4"
)

// Handler ...
type Handler struct {
	services *service.Service
}

// NewHandler ...
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// Init ...
func (h *Handler) Init(e *echo.Echo) {

	g := e.Group("/api/v1")

	g.GET("/posts", h.getAllPosts)
	g.POST("/posts", h.createPost)
	g.GET("/posts/:id", h.getOnePost)
	g.PUT("/posts/:id", h.updatePost)
	g.DELETE("/posts/:id", h.deletePost)

	g.GET("/comments", h.getAllComments)
	g.POST("/comments", h.createComment)
	g.GET("/comments/:id", h.getOneComment)
	g.PUT("/comments/:id", h.updateComment)
	g.DELETE("/comments/:id", h.deleteComment)

}
