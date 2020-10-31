package api

import (
	"github.com/labstack/echo"
	"net/http"
)

//func GetUsers() echo.HandlerFunc {
//	return func(c echo.Context) (err error) {
//		return c.String(http.StatusOK, "Hello, World!")
//	}
//}

func GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
