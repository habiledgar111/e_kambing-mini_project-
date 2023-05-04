package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mini_project/config"
	"mini_project/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtcustomclaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

var (
// DB   *gorm.DB
// user *model.User
)

func init() {
	//connect db
	// dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// DB = db
	// DB.AutoMigrate(user)
}

func createJWT(user model.User) interface{} {
	email := user.Email
	pass := user.Password
	claims := &jwtcustomclaims{
		email,
		pass,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := temp.SignedString([]byte("secret"))

	if err != nil {
		return err.Error()
	}

	return token
}

func Login(c echo.Context) error {
	var users model.User
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	email := fmt.Sprintf("%v", json_map["email"])
	password := fmt.Sprintf("%v", json_map["password"])

	// user = model.GetUserFromEmail(email)
	if err := config.DB.Where("email = @email", sql.Named("email", email)).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if users.Password != password {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "wrong password",
		})
	}
	token := createJWT(users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}

func CreateUser(c echo.Context) error {
	var users model.User
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	name := fmt.Sprintf("%v", json_map["name"])
	email := fmt.Sprintf("%v", json_map["email"])
	password := fmt.Sprintf("%v", json_map["password"])
	users = model.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	result := config.DB.Create(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create user",
		"rowaffected": result.RowsAffected,
		"querybool":   result.QueryFields,
	})
}
