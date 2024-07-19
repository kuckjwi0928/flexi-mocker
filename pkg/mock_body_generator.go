package pkg

import (
	"github.com/kuckjwi0928/flexi-mocker/configs"
)

const bindingChar = '$'

type MockBodyGenerator struct {
}

func NewMockBodyGenerator() *MockBodyGenerator {
	return &MockBodyGenerator{}
}

func (m *MockBodyGenerator) GetBody(route configs.RouteConfig, parameter map[string]any) any {
	return m.generate(route.Response, parameter)
}

func (m *MockBodyGenerator) generate(response any, parameter map[string]any) any {
	switch response.(type) {
	case string:
		return m.resolve(response.(string), parameter)
	case map[string]any:
		body := make(map[string]any)
		for k, v := range response.(map[string]any) {
			body[k] = m.generate(v, parameter)
		}
		return body
	case []any:
		body := make([]any, 0)
		for _, v := range response.([]any) {
			body = append(body, m.generate(v, parameter))
		}
		return body
	default:
		return response
	}
}

func (m *MockBodyGenerator) resolve(origin string, parameter map[string]any) any {
	if origin[0] == bindingChar {
		if value, ok := parameter[origin[1:]]; ok {
			return value
		}
	}
	return origin
}
