package cli

import (
	"testing"

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
		assert.Equal(t, tc.expected, ParseKVFlagValue(tc.value))
	}
}