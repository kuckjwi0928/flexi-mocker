package server

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/kuckjwi0928/flexi-mocker/configs"
)

type Server struct {
	app            *fiber.App
	config         *configs.Config
	mockingHandler *MockingHandler
}

func NewServer(config *configs.Config) *Server {
	return &Server{
		app:            fiber.New(),
		config:         config,
		mockingHandler: newMockingHandler(),
	}
}

func (s *Server) setupRoutes() {
	for _, route := range s.config.Routes {
		r := route
		s.app.Add([]string{r.Method}, r.Path, func(ctx fiber.Ctx) error {
			return s.mockingHandler.handle(ctx, r)
		})
	}
}

func (s *Server) Run(port int) {
	s.setupRoutes()
	err := s.app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
