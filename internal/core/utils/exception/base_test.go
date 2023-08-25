package exception

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppErrorError(t *testing.T) {
	tests := []struct {
		name   string
		appErr AppError
		want   string
	}{
		{
			name:   "Test error message",
			appErr: AppError{Type: 1, Message: "This is an error"},
			want:   "This is an error",
		},
		// Add more test cases here
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.appErr.Error()
			assert.Equal(t, test.want, got)
		})
	}
}
