package restweb

import (
	"html/template"
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
