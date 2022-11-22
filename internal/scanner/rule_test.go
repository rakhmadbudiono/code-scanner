package scanner_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rakhmadbudiono/code-scanner/internal/scanner"
)

func TestGetRules(t *testing.T) {
	cases := []struct {
		expected []scanner.Rule
	}{
		{
			expected: []scanner.Rule{
				{
					ID:          "G400",
					Type:        "sast",
					Severity:    scanner.High,
					Description: "private key exposed",
					Checker: func(line string) bool {
						return strings.Contains(line, "private_key")
					},
				},
				{
					ID:          "G401",
					Type:        "sast",
					Severity:    scanner.High,
					Description: "public key exposed",
					Checker: func(line string) bool {
						return strings.Contains(line, "public_key")
					},
				},
			},
		},
	}

	for _, tc := range cases {
		rules := scanner.GetRules()

		for i, rule := range rules {
			assert.Equal(t, rule.Checker("test"), tc.expected[i].Checker("test"))
			tc.expected[i].Checker = nil
			rules[i].Checker = nil
		}

		assert.Equal(t, tc.expected, rules)
	}
}
