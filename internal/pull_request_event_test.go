package internal_test

import (
	"testing"

	"github.com/kazuki-iwanaga/pr2trace/internal"
)

func TestPullRequestEventName(t *testing.T) {
	t.Parallel()
	testcases := []struct { // nolint:wsl
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
		e := internal.NewPullRequestEvent(tt.name)

		if e.Name() != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, e.Name())
		}
	}
}
