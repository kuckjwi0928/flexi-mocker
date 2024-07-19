package pkg

import (
	"reflect"
	"testing"
)

func TestMockBodyGenerator_resolve(t *testing.T) {
	type args struct {
		origin    string
		parameter map[string]any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Resolved header binding parameter",
			args: args{
				origin:    "$header.x-custom-value",
				parameter: map[string]any{"header.x-custom-value": "Hello World!"},
			},
			want: "Hello World!",
		},
		{
			name: "Resolved query binding parameter",
			args: args{
				origin:    "$query.value",
				parameter: map[string]any{"query.value": "Hello World!"},
			},
			want: "Hello World!",
		},
		{
			name: "Resolved path binding parameter",
			args: args{
				origin:    "$path.id",
				parameter: map[string]any{"path.id": "Hello World!"},
			},
			want: "Hello World!",
		},
		{
			name: "Resolved body binding parameter",
			args: args{
				origin:    "$body.value",
				parameter: map[string]any{"body.value": "Hello World!"},
			},
			want: "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockBodyGenerator{}
			if got := m.resolve(tt.args.origin, tt.args.parameter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
