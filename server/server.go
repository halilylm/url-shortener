package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/halilylm/url-shortner/config"
	"github.com/labstack/echo/v4"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

type Server struct {
	cfg     *config.Config
	e       *echo.Echo
	address string
}

func New(cfg *config.Config) *Server {
	address := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	e := echo.New()
	e.Server.WriteTimeout = cfg.WriteTimeout
	e.Server.ReadTimeout = cfg.ReadTimeout
	e.Server.IdleTimeout = cfg.IdleTimeout
	e.HideBanner = true
	e.Server.ReadHeaderTimeout = cfg.ReadHeaderTimeout
	return &Server{
		cfg:     cfg,
		e:       e,
		address: address,
	}
}

func (s *Server) Run() {
	s.setupRoutes()
	fmt.Println("start listening on " + s.address)
	go func() {
		if err := s.e.Start(s.address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()
	fmt.Println("here")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	fmt.Println("gracefully shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.GracefulShutdownTimeout)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("shut down!")
}
