package converter

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// StrToTime 字符串转时间
//
// @param str 时间字符串，示例格式为 "2006-01-02 15:04:05.999999"
//
// @return int64 转换后的时间戳，参数错误时返回 0 值
func StrToTime(timeStr string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05.999999", timeStr, time.Local)

	if err != nil {
		logrus.Errorf("解析时间字符串失败: %v", err)
		return time.Time{}
	}

	return t
}
