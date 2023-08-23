package exception

import (
	"testing"
)

func TestAppError_Error(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.appErr.Error(); got != tt.want {
				t.Errorf("AppError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
