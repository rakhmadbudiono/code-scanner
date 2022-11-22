package scanner_test

import (
	"errors"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/rakhmadbudiono/code-scanner/internal/scanner"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	cases := []struct {
		input    string
		expected []scanner.Finding
		err      error
	}{
		{
			input:    "https://github.com/rakhmadbudiono/duck-pic-service",
			expected: []scanner.Finding{},
			err:      nil,
		},
		{
			input:    "",
			expected: nil,
			err:      errors.New("processing message, error cloning: URL field is required ()"),
		},
	}

	for _, tc := range cases {
		sp := &scanner.ScanProcessor{FileSystem: memfs.New()}
		findings, err := sp.Process(tc.input)

		assert.Equal(t, tc.expected, findings)
		assert.Equal(t, tc.err, err)
	}
}
