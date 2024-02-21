package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kuckjwi0928/flexi-mocker/configs"
	"github.com/kuckjwi0928/flexi-mocker/pkg"
	"net/http"
)

type MockingHandler struct {
	bodyGenerator *pkg.MockBodyGenerator
}

func newMockingHandler() *MockingHandler {
	return &MockingHandler{
		bodyGenerator: pkg.NewMockBodyGenerator(),
	}
}

func (m *MockingHandler) handle(c fiber.Ctx, route configs.RouteConfig) error {
	return c.JSON(m.bodyGenerator.GetBody(route, m.extractRequest(c)))
}

func (m *MockingHandler) extractRequest(c fiber.Ctx) map[string]any {
	_m := make(map[string]any)
	for _, v := range c.Route().Params {
		_m[v] = c.Params(v)
	}
	if c.Method() == http.MethodGet {
		for k, v := range c.Queries() {
			_m[k] = v
		}
	}
	// TODO: extract request body for post request
	return _m
}
