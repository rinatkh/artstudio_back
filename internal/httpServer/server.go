package httpServer

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/pkg/httpErrorHandler"
	"github.com/sirupsen/logrus"
)

// Server struct
type Server struct {
	fiber *fiber.App
	cfg   *config.Config
	log   *logrus.Entry
}

func NewServer(cfg *config.Config, handler *httpErrorHandler.HttpErrorHandler, log *logrus.Entry) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{ErrorHandler: handler.Handler, DisableStartupMessage: true}),
		cfg:   cfg,
		log:   log,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.fiber); err != nil {
		s.log.Fatalln("Cannot map handlers: ", err)
	}
	s.log.Println(fmt.Sprintf("Start server on %s:%s", s.cfg.Server.Host, s.cfg.Server.Port))
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		s.log.Fatalln("Error starting Server: ", err)
	}

	return nil
}
