package converter

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInt64ToStr(t *testing.T) {
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
			got := Int64ToStr(tt.input)
			assert.Equal(t, tt.expected, got, "Int64ToString(%v) = %v; 期望值 %v", tt.input, got, tt.expected)
		})
	}
}

func TestStrToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			name:     "2023-12-25 15:04:05.123456",
			input:    "2023-12-25 15:04:05.123456",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 123456000, time.Local),
		},
		{
			name:     "2025-04-20 17:00:36.311701",
			input:    "2025-04-20 17:00:36.311701",
			expected: time.Date(2025, 4, 20, 17, 0, 36, 311701000, time.Local),
		},
		{
			name:     "2000-01-01 00:00:00.000000",
			input:    "2000-01-01 00:00:00.000000",
			expected: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "错误格式",
			input:    "invalid-time-format",
			expected: time.Time{},
		},
		{
			name:     "空字符串",
			input:    "",
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrToTime(tt.input)
			assert.Equal(t, tt.expected, got, "StrToTime(%v) = %v; 期望值 %v", tt.input, got, tt.expected)
		})
	}
}
