package restweb

import (
	"container/list"
	"net/http"
	"reflect"
)

// GET（SELECT）：从服务器取出资源（一项或多项）。
// POST（CREATE）：在服务器新建一个资源。
// PUT（UPDATE）：在服务器更新资源（客户端提供改变后的完整资源）。
// PATCH（UPDATE）：在服务器更新资源（客户端提供改变的属性）。
// DELETE（DELETE）：从服务器删除资源。
// HEAD：获取资源的元数据。
// OPTIONS：获取信息，关于资源的哪些属性是客户端可以改变的。

type Router interface {
	Post()
	Get()
	Delete()
	Put()
	Patch()
	Head()
	Options()
}

func CallMethod(c interface{}, m string, rv []reflect.Value) {
	rc := reflect.ValueOf(c)
	rm := rc.MethodByName(m)
	rm.Call(rv)
}

func GetReflectValue(i ...interface{}) (rv []reflect.Value) {
	for _, j := range i {
		rw := reflect.ValueOf(j)
		rv = append(rv, rw)
	}
	return
}

var routerList = &list.List{}

func AddRouter(method string, pattern string, controllerName string, action string) {

	for e := controllerList.Front(); e != nil; e = e.Next() {
		c := e.Value.(reflect.Type)
		Logger.Debug(c.Name())

		if c.Name() == controllerName {

			routerList.PushBack(
				&Control{Type: c, Method: method,
					Pattern: pattern, Action: action})
			break
		}
	}
}

var FileMap = map[string]http.Handler{}

//添加静态文件路由
func AddFile(pattern string, fileHandler http.Handler) {
	FileMap[pattern] = fileHandler
}
