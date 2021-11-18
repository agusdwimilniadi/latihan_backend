package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	// routing
	e.GET("/", HelloController)
	e.GET("/news", GetNewsController)
	e.GET("/news/:newsId", GetDetailNewsController)
	e.GET("/login", LoginController)

	e.Start(":8000")
}

type News struct {
	Id      int
	Title   string
	Content string
	Author  string
}
type User struct {
	Email    string `json:"email"`
	Password string `password:"password"`
}
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func GetDetailNewsController(c echo.Context) error {
	newsId, _ := strconv.Atoi(c.Param("newsId"))

	var data = News{1, "A", "A", "A"}

	var response = BaseResponse{
		Code:    newsId,
		Message: "Succes",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func GetNewsController(c echo.Context) error {
	lokasi := c.QueryParam("lokasi")
	var data = []News{
		{1, "A", "A", "A"},
		{2, "B", "B", "B"},
	}
	var response = BaseResponse{
		Code:    http.StatusOK,
		Message: lokasi,
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginController(e echo.Context) error {
	// email := e.FormValue("email")
	// password := e.FormValue("password")
	var user User
	e.Bind(&user)

	if user.Password == "" {
		return e.JSON(http.StatusBadRequest, BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Kosong",
			Data:    nil,
		})
	}
	// user := User{
	// 	Email:    email,
	// 	Password: password,
	// }
	return e.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Succes",
		Data:    user,
	})
}

func ConnectDB() {
	dsn := "root:123edsaq1@tcp(127.0.0.1:3307)/kel_d?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database not connect")
	}
}
