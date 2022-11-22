package scanner_test

import (
	"testing"

	"github.com/rakhmadbudiono/code-scanner/internal/scanner"
	"github.com/stretchr/testify/assert"
)

func TestNewFinding(t *testing.T) {
	cases := []struct {
		input struct {
			path string
			line int
			rule scanner.Rule
		}
		expected *scanner.Finding
	}{
		{
			input: struct {
				path string
				line int
				rule scanner.Rule
			}{
				path: "test.txt",
				line: 1,
				rule: scanner.Rule{
					ID:          "test0",
					Type:        "test",
					Severity:    scanner.Low,
					Description: "test",
					Checker: func(s string) bool {
						return true
					},
				},
			},
			expected: &scanner.Finding{
				Type:   "test",
				RuleID: "test0",
				Location: scanner.Location{
					Path: "test.txt",
					Positions: scanner.Positions{
						Begin: scanner.Begin{
							Line: 1,
						},
					},
				},
				Metadata: scanner.Metadata{
					Description: "test",
					Severity:    scanner.Low,
				},
			},
		},
	}

	for _, tc := range cases {
		finding := scanner.NewFinding(tc.input.path, tc.input.line, tc.input.rule)

		assert.Equal(t, tc.expected, finding)
	}
}
