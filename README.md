restweb
=======

light web framework for go

##Features

1. 支持路由配置文件, 支持控制器方法参数传入
2. 模板自动渲染
3. session 管理
4. 表单验证
5. 过滤和拦截
	

##TODO List:

4. 配置文件

##使用文档

###简介

restweb 采用 MCV 模式并支持RESTful API 设计，是一个具有丰富特性的轻量级web框架

### Controller

#### 简介

以下是restweb自带的Controller

	type Controller struct {
		*Context 
		Action string //method of controller being callled
		Name   string
	}

使用自定义控制器时，应该嵌入restweb.Controller

	type MainController struct {
		restweb.Controller
	}
	
注册控制器，注册的控制器必须具有 Router 接口
	
	func RegisterController(controller Router)
####路由

1. 可以将路由定义在一个单独的文件中(config/router.conf)
	
	路由定义规则是:
	
		(METHOD) (URL Pattern) (Controller.Action)
	路由优先级按文件中的路由顺序、URL支持正则表达式匹配
	
	一个简单的样例：
		
		#comment
		#METHOD URI				CONTROLLER_ACTION
		GET 	/ 				HomeController.Index
		GET 	^/users/(\w+)	UserController.Detail
		POST 	^/users/\w+		UserController.Update
2. 支持使用添加路由函数直接注册路由

		AddRouter(method string, pattern string, controllerName string, action string)
		
####上下文

上下文Context定义为：

	type Context struct {
		R      *http.Request
		W      http.ResponseWriter
		Input  url.Values
		Output map[string]interface{}
	}
1. Input  

	内部包括表单数据和URL参数，可以使用的方法有Get、Set、Add、Encode、Del，可以直接用Input["Para"]方式获得输入参数
2. OutPut  

	保存数据，用于模板的渲染
3. 设置Session值

		func (c *Context) SetSession(key string, value string) 
4. 获取Session值

		func (c *Context) GetSession(key string) string 
5. 销毁Session

		func (c *Context) DeleteSession()
6. 重定向

		func (c *Context) Redirect(urlStr string, code int)
7. http错误

		func (c *Context) Error(err string, code int)
		
###过滤器
过滤器是可以针对特定路由和上下文环境处理的restweb中间件

	type Filter func(ctx *Context) bool

1. 过滤器函数返回值为true则为拦截，支持控制器方法调用前和后拦截、过滤
2. 过滤器按注册的顺序安排优先级，注册早的优先级高
3. 对于一个url，如果一个拦截器被执行，将立即停止执行其后的过滤器和控制器方法
4. 注册过滤器

		func RegisterFilters(method string, pattern string, when int, filter Filter) 