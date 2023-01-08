package sicepat

import (
	"context"
	"sicepat/internal/api"
	"sicepat/internal/storage"

	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultReadTimeout  = 3 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func New(cfg Config) (*Server, error) {
	stopCh := make(chan struct{})

	return NewServer(cfg, stopCh)
}

func NewServer(cfg Config, stopCh chan struct{}) (*Server, error) {

	storage, err := storage.NewStorage(cfg)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	// build server
	server := &Server{
		stopCh: stopCh,
		httpServer: &http.Server{
			Addr:         cfg.ServerAddress(),
			Handler:      api.BuildRouter(storage),
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		},
	}

	return server, nil
}

type Server struct {
	stopCh chan struct{}

	httpServer *http.Server
}

func (s *Server) Address() string {
	return s.httpServer.Addr
}

func (s *Server) Start() {
	logrus.WithField("address", s.httpServer.Addr).Info("starting HTTP server")
	_ = s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		wg.Done()

		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.Warn("failed to shutdown HTTP server", logrus.Fields{"error": err})
		}
	}()

	wg.Wait()
}

type Config interface {
	// HTTP server
	ServerAddress() string
	storage.Config
}
