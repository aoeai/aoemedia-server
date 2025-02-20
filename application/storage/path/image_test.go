package path

import (
	"fmt"
	"testing"
	"time"
)

func TestYearMonth(t *testing.T) {
	type args struct {
		createTime time.Time
		expected   string
	}

	tests := []args{
		{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), "2021-01"},
		{time.Date(2021, 12, 31, 23, 59, 59, 0, time.UTC), "2021-12"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("使用年月作为文件夹：%v->%v", tt.createTime, tt.expected)
		t.Run(name, func(t *testing.T) {
			if yearMonth := YearMonthOf(tt.createTime); yearMonth != tt.expected {
				t.Errorf("实际值 = %v, 期望值 %v", yearMonth, tt.expected)
			}
		})
	}
}
