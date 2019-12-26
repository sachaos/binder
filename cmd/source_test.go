package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSource(t *testing.T) {
	testCases := []struct{argument string; expected []*Source}{
		{
			argument: "app:yes",
			expected: []*Source{
				{
					Name: "app",
					Command: "yes",
				},
			},
		},
		{
			argument: "app : yes",
			expected: []*Source{
				{
					Name: "app",
					Command: "yes",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.argument, func(t *testing.T) {
			sources, err := ParseSource(strings.NewReader(testCase.argument))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, sources, testCase.expected)
		})
	}
}
