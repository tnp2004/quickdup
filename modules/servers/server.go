package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tnp2004/quickdup/configs"
	"github.com/tnp2004/quickdup/pkg/databases"
)

type Server struct {
	cfg    *configs.Config
	server *echo.Echo
	db     databases.Database
}

func NewServer(cfg *configs.Config, db databases.Database) *Server {
	e := echo.New()
	return &Server{
		cfg:    cfg,
		server: e,
		db:     db,
	}
}

func (s *Server) Start() {
	s.registerModuleRouters()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go s.listenAndServe()

	s.gracefulShutdown(ctx)
}

func (s *Server) listenAndServe() {
	log.Printf("start the server on port %s", s.cfg.Server.Port)

	address := fmt.Sprintf(":%s", s.cfg.Server.Port)
	if err := s.server.Start(address); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error server is shutting down. Err: %s", err.Error())
	}
}

func (s *Server) gracefulShutdown(ctx context.Context) {
	<-ctx.Done()
	log.Println("gracefully shutdown the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.server.Logger.Fatal(err)
	}
}
