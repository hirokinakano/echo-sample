package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"

	"echo-sample/api"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Logger.Debug()

	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())

	v1 := e.Group("/api/v1")
	{
		v1.GET("/users", api.GetUsers)
	}

	return e
}
