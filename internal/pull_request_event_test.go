package internal_test

import (
	"testing"

	"github.com/kazuki-iwanaga/pr2trace/internal"
	"github.com/stretchr/testify/assert"
)

func TestPullRequestEventName(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		expected string
	}{
		{
			name:     "test",
			expected: "test",
		},
		{
			name:     "test2",
			expected: "test2",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := internal.NewPullRequestEvent(tt.name)

			actual := e.Name()
			if e.Name() != tt.expected {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
