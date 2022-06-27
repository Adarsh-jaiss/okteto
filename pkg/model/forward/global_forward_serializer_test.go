package forward

import (
	"reflect"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestGlobalForward_MarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		data     GlobalForward
	}{
		{
			name:     "service-with-port",
			expected: "8080:svc:5214",
			data:     GlobalForward{Local: 8080, Remote: 5214, ServiceName: "svc"},
		},
		{
			name:     "service-with-port",
			expected: "27017:mongodb:27017",
			data:     GlobalForward{Local: 27017, Remote: 27017, ServiceName: "mongodb"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := yaml.Marshal(tt.data)
			if err != nil {
				t.Error(err)
			}

			outStr := strings.Trim(string(b), "\n")
			if outStr != tt.expected {
				t.Errorf("didn't marshal correctly. Actual '%+v', Expected '%+v'", outStr, tt.expected)
			}

		})
	}
}

func TestGlobalForward_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		expected  GlobalForward
		expectErr bool
	}{
		{
			name:      "basic",
			data:      "8080:9090",
			expectErr: true,
		},
		{
			name:      "name of service not specified",
			data:      "8080::9090",
			expectErr: true,
		},
		{
			name:      "equal",
			data:      "8080:8080",
			expectErr: true,
		},
		{
			name:      "service-with-port",
			data:      "8080:svc:5214",
			expectErr: false,
			expected:  GlobalForward{Local: 8080, Remote: 5214, ServiceName: "svc"},
		},
		{
			name:      "bad-local-port",
			data:      "local:8080",
			expectErr: true,
		},
		{
			name:      "service-with-bad-port",
			data:      "8080:svc:bar",
			expectErr: true,
		},
		{
			name:      "too-little-parts",
			data:      "8080",
			expectErr: true,
		},
		{
			name:      "too-many-parts",
			data:      "8080:svc:8082:8019",
			expectErr: true,
		},
		{
			name:      "service-at-end",
			data:      "8080:8081:svc",
			expectErr: true,
		},
		{
			name:      "just-service",
			data:      "8080:svc",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result GlobalForward
			if err := yaml.Unmarshal([]byte(tt.data), &result); err != nil {
				if tt.expectErr {
					return
				}

				t.Fatal(err)
			}

			if tt.expectErr {
				t.Fatal("didn't got expected error")
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("didn't unmarshal correctly. Actual '%+v', Expected '%+v'", result, tt.expected)
			}

			out, err := yaml.Marshal(result)
			if err != nil {
				t.Fatal(err)
			}

			outStr := string(out)
			outStr = strings.TrimSuffix(outStr, "\n")

			if !reflect.DeepEqual(outStr, tt.data) {
				t.Errorf("didn't unmarshal correctly. Actual '%+v', Expected '%+v'", outStr, tt.data)
			}
		})
	}
}

func TestGlobalForwardExtended_MarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		data     GlobalForward
	}{
		{
			name:     "service-name",
			expected: "8080:svc:9090",
			data:     GlobalForward{Local: 8080, Remote: 9090, ServiceName: "svc"},
		},
		{
			name:     "service-name-and-labels",
			expected: "8080:svc:5214",
			data:     GlobalForward{Local: 8080, Remote: 5214, ServiceName: "svc", Labels: map[string]string{"key": "value"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := yaml.Marshal(tt.data)
			if err != nil {
				t.Error(err)
			}

			outStr := strings.Trim(string(b), "\n")
			if outStr != tt.expected {
				t.Errorf("didn't marshal correctly. Actual '%+v', Expected '%+v'", outStr, tt.expected)
			}

		})
	}
}
