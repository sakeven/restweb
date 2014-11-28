package restweb

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

var funcMap map[string]interface{}

// AddFuncMap add new template function.
func AddFuncMap(key string, f interface{}) {
	funcMap[key] = f
}

//ParseFiles 合成新的tpl文件.
func ParseFiles(tplfiles ...string) (*template.Template, error) {
	t := template.New("layout.tpl").Funcs(template.FuncMap(funcMap))
	t, err := t.ParseFiles(tplfiles...)
	return t, err
}

// initFuncMap init FuncMap.
func initFuncMap() {
	funcMap = make(map[string]interface{})
	funcMap["NumAdd"] = NumAdd
	funcMap["NumSub"] = NumSub
	funcMap["ShowTime"] = ShowTime
	funcMap["ShowGapTime"] = ShowGapTime
}

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

func GetTime() (t int64) {
	t = time.Now().Unix()
	return
}

func GetAction(path string, pos int) string {
	path = strings.Trim(path, "/")
	pathsplit := strings.Split(path, "/")
	if pos >= 0 && pos < len(pathsplit) {
		return pathsplit[pos]
	}
	return ""
}
