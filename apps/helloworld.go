package apps

import (
	"github.com/labstack/echo"
	"net/http"
)

//HelloWorld 一个简单的 HTTP 请求
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
