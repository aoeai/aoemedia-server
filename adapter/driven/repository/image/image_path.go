package image

import "time"

func YearMonthOf(createTime time.Time) string {
	return createTime.Format("2006-01")
}
