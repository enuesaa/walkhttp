package cli

import (
	"testing"

	"github.com/enuesaa/walkin/internal/web"
	"github.com/stretchr/testify/assert"
)

func TestParseKVFlagValue(t *testing.T) {
	cases := []struct {
		value    string
		expected map[string]string
	}{
		{
			value:    "a=b",
			expected: map[string]string{"a": "b"},
		},
		{
			value:    "a=b,c=d",
			expected: map[string]string{"a": "b", "c": "d"},
		},
	}

	for _, tc := range cases {
		ret, _ := ParseKVFlagValue(tc.value)
		assert.Equal(t, tc.expected, ret)		
	}
}

func TestParseFlagsToServeConfig(t *testing.T) {
	cases := []struct {
		readLocalFilesFlag string
		proxyFlag          string
		expected           web.ServeConfig
	}{
		{
			readLocalFilesFlag: "path=/*",
			proxyFlag: "path=/aaa/*,url=https://example.com",
			expected: web.ServeConfig{
				Paths: map[string]web.Behavior{
					"/*": {
						Behavior: web.ReadLocalFiles,
					},
					"/aaa/*": {
						Behavior: web.Proxy,
						ProxyConfig: web.ProxyConfig{
							Url: "https://example.com",
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		serveConfig, _ := ParseFlagsToServeConfig(tc.readLocalFilesFlag, tc.proxyFlag)
		assert.Equal(t, tc.expected, serveConfig)		
	}
}
