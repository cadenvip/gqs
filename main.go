package main

import (
	"github.com/labstack/echo"
	"github.com/sunisdown/gqs/apps"
)

func main() {
	e := echo.New()
	e.Static("/", "public")
	// 可以使用 Group 来区分不同的 App，比如这里的 HelloWorld 是一个单独的 App，就可以用 Group 来做隔离。
	g := e.Group("/helloworld")
	g.GET("/", apps.HelloWorld)

	// 可以返回 JSON
	j := e.Group("/jsonp")
	j.GET("", apps.JSONP)

	// 正常的分路由
	e.GET("/grpc", apps.GRPCDEMO)
	e.Logger.Fatal(e.Start(":1323"))
}
