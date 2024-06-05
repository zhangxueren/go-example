package utils

import "time"

var Time = new(t)

type t struct {
}

// 获取当前时间戳
func (t) Now() int {
	return int(time.Now().Unix())
}

func (t) NowMs() int64 {
	return time.Now().UnixNano() / 1e6
}

func (t) Format(timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

/**
 * @Description: 获取今天日期
 * @return string 2021-01-01
 */
func (t) Today() string {
	return time.Now().Format("2006-01-02")
}

func (t) TodayUnixTimeStamp() int64 {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return today.Unix()
}

func (t) RemainingSecondsOfToday() uint64 {
	now := time.Now()
	year, month, day := now.Date()
	endOfDay := time.Date(year, month, day, 23, 59, 59, 0, now.Location())
	remaining := endOfDay.Sub(now)
	return uint64(remaining.Seconds())
}
