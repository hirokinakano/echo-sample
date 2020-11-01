package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"

	"echo-sample/db"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(c echo.Context) error {
	db := db.GormConnect()

	result := db.Find(&[]User{})
	return c.JSON(http.StatusOK, result)
}
