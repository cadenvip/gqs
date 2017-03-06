package apps

import (
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

func JSONP(c echo.Context) error {
	callback := c.QueryParam("callback")
	var context struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}
	context.Response = "Sent via JSONP"
	context.Timestamp = time.Now().UTC()
	context.Random = rand.Intn(1000)
	return c.JSONP(http.StatusOK, callback, &context)
}
