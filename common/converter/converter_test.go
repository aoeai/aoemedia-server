package converter

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64ToString(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{
			name:     "正常整数",
			input:    42,
			expected: "42",
		},
		{
			name:     "负数",
			input:    -123,
			expected: "-123",
		},
		{
			name:     "零",
			input:    0,
			expected: "0",
		},
		{
			name:     "int64最大值",
			input:    math.MaxInt64,
			expected: "9223372036854775807",
		},
		{
			name:     "int64最小值",
			input:    math.MinInt64,
			expected: "-9223372036854775808",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64ToString(tt.input)
			assert.Equal(t, tt.expected, got, "Int64ToString(%v) = %v; 期望值 %v", tt.input, got, tt.expected)
		})
	}
}
