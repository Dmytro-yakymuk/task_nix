package handler

import (
	"net/http"

	"github.com/Dmytro-yakymuk/task_nix/internal/service"
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
// to present data in xml format, add to the address `?type=xml`
func (h *Handler) Init() {

	http.HandleFunc("/api/v1/posts", h.routePosts)  // GET, POST
	http.HandleFunc("/api/v1/posts/", h.routePosts) // GET, PUT, DELETE

}
