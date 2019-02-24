package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/", WebRoot)

	// 加载templates文件夹下所有的文件
	// engine.LoadHTMLGlob("templates/**/*")
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// engine.LoadHTMLGlob("templates/*")
	// engine.LoadHTMLGlob("templates/*/*.html")
	engine.LoadHTMLGlob("templates/*")

	engine.Static("static", "./static")
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 注册一个动态路由
	// 可以匹配 /user/joy
	// 不能匹配 /user 和 /user/
	engine.GET("/user/:name", func(c *gin.Context) {
		// 使用 c.Param(key) 获取 url 参数
		// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
		c.HTML(http.StatusOK, "layui.html", gin.H{
			"title": "GIN: 测试加载HTML模板",
		})
	})

	// 注册一个高级的动态路由
	// 该路由会匹配 /user/john/ 和 /user/john/send
	// 如果没有任何路由匹配到 /user/john, 那么他就会重定向到 /user/john/，从而被该方法匹配到
	engine.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 注册一个动态路由
	// 可以匹配 /user/joy
	// 不能匹配 /user 和 /user/
	engine.GET("/demo/:name/*acton", func(c *gin.Context) {
		// 使用 c.Param(key) 获取 url 参数
		// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
		name := c.Param("name")
		log.Println("name==", name)
		c.HTML(http.StatusOK, name, gin.H{
			"title": "GIN: 测试加载HTML模板",
		})
	})

	// log.Println("11111111111111111111111", strings.Compare(string(1), "1"))

	// 绑定端口，然后启动应用
	engine.Run(":9100")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	// context.String(http.StatusOK, "hello, world")
}
