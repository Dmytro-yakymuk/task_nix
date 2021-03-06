package server

import (
	"net/http"

	"github.com/Dmytro-yakymuk/task_nix/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20, // 1 MB
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// func (s *Server) Stop(ctx context.Context) error {
// 	return s.httpServer.Shutdown(ctx)
// }
