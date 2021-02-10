package handler

import (
	"github.com/Dmytro-yakymuk/task_nix/internal/service"

	"github.com/labstack/echo/v4"

	_ "github.com/Dmytro-yakymuk/task_nix/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
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

func (h *Handler) Init(e *echo.Echo) {

	e.GET("/auth", h.auth)
	e.GET("/auth/callback", h.authCallback)
	e.GET("/logout", h.logout)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1")

	g.GET("/posts", h.getAllPosts)
	g.GET("/posts/:id", h.getOnePost)

	g.GET("/comments", h.getAllComments)
	g.GET("/comments/:id", h.getOneComment)

	au := g.Group("")
	au.Use(h.userIdentity)

	au.POST("/posts", h.createPost)
	au.PUT("/posts/:id", h.updatePost)
	au.DELETE("/posts/:id", h.deletePost)

	au.POST("/comments", h.createComment)
	au.PUT("/comments/:id", h.updateComment)
	au.DELETE("/comments/:id", h.deleteComment)

}
