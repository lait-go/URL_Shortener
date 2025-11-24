package http_server

import (
	"context"
	"net/http"
	"time"
)

type Config struct {
	ShutdownTimeout time.Duration `default:"10s" envconfig:"SHUTDOWN_TIMEOUT"`
	Port            string        `default:":8081" envconfig:"HTTP_PORT"`
	Swagger         string        `envconfig:"SWAG_URL"`
}

type Server struct {
	HTTPServer *http.Server
}

func New(r http.Handler, c Config) *Server {
	r = http.TimeoutHandler(r, time.Second*5, "request timeout")

	h := &Server{
		HTTPServer: &http.Server{
			Addr:    c.Port,
			Handler: r,
		},
	}

	return h
}

func (s *Server) Run() error {
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Close(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    return s.HTTPServer.Shutdown(ctx)
}
