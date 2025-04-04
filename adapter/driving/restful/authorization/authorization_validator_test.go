package authorization

import "testing"

func TestAuthorization_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		auth     Authorization
		expected bool
	}{
		{
			name: "should return true when Missing is true",
			auth: Authorization{
				Missing: true,
				UserId:  100,
			},
			expected: true,
		},
		{
			name: "should return true when UserId is less than 1",
			auth: Authorization{
				Missing: false,
				UserId:  0,
			},
			expected: true,
		},
		{
			name: "should return false when Missing is false and UserId is valid",
			auth: Authorization{
				Missing: false,
				UserId:  100,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.auth.Invalid(); got != tt.expected {
				t.Errorf("Authorization.Invalid() = %v, want %v", got, tt.expected)
			}
		})
	}
}
