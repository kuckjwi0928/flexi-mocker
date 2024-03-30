package pkg

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/kuckjwi0928/flexi-mocker/configs"
	"net/http"
)

const (
	pathKey   = "path"
	headerKey = "header"
	queryKey  = "query"
	bodyKey   = "body"
)

type MockingHandler struct {
	bodyGenerator *MockBodyGenerator
}

func newMockingHandler() *MockingHandler {
	return &MockingHandler{
		bodyGenerator: NewMockBodyGenerator(),
	}
}

func (m *MockingHandler) handle(c fiber.Ctx, route configs.RouteConfig) error {
	body, err := m.extractRequest(c)
	if err != nil {
		return err
	}
	return c.JSON(m.bodyGenerator.GetBody(route, body))
}

func (m *MockingHandler) extractRequest(c fiber.Ctx) (map[string]any, error) {
	_m := make(map[string]any)
	for k, v := range c.GetReqHeaders() {
		arr := make([]string, len(v))
		for i, v := range v {
			arr[i] = v
		}
		_m[fmt.Sprintf("%s.%s", headerKey, k)] = arr
	}
	for _, v := range c.Route().Params {
		_m[fmt.Sprintf("%s.%s", pathKey, v)] = c.Params(v)
	}
	switch c.Method() {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		if c.Body()[0] == '{' {
			body := make(map[string]any)
			err := c.Bind().JSON(&body)
			if err != nil {
				return nil, err
			}
			for k, v := range body {
				_m[fmt.Sprintf("%s.%s", bodyKey, k)] = v
			}
		} else {
			bodies := make([]any, 0)
			err := c.Bind().JSON(&bodies)
			if err != nil {
				return nil, err
			}
			for _, body := range bodies {
				for k, v := range body.(map[string]any) {
					_m[fmt.Sprintf("%s.%s", bodyKey, k)] = v
				}
			}
		}
	case http.MethodGet:
		for k, v := range c.Queries() {
			_m[fmt.Sprintf("%s.%s", queryKey, k)] = v
		}
	}
	return _m, nil
}
