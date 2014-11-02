package goweb

import (
	"fmt"
	"regexp"
	"time"
)

// ShowNext 返回num加1的值
func ShowNext(num int) (next int) {
	next = num + 1
	return
}

// ShowTime 将unixtime转换为当地时间
func ShowTime(unixtime int64) string {
	return time.Unix(unixtime, 0).Local().Format("2006-01-02 15:04:05")
}

// NumAdd 将两数相加
func NumAdd(a int, b int) (ret int) {
	ret = a + b
	return
}

// NumSub 两数相减a-b
func NumSub(a int, b int) (ret int) {
	ret = a - b
	return
}

// 格式化间隔时间
func ShowGapTime(gaptime int64) string {
	sec := gaptime % 60
	hour := gaptime / 3600
	minute := (gaptime - hour*3600) / 60
	return fmt.Sprintf("%d:%02d:%02d", hour, minute, sec)
}

// 检查邮箱地址格式
func CheckMail(mail string) bool {
	matchStr := `^\w(\.?\w)*@\w(\.?\w)*\.[A-Za-z]+$`
	tagRx := regexp.MustCompile(matchStr)
	tagString := tagRx.FindString(mail)
	if tagString == mail {
		return true
	}
	return false
}
