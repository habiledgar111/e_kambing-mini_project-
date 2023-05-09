package main

import (
	"mini_project/config"
	"mini_project/controller"
	"mini_project/model"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

var (
	kambing   *model.Kambing
	user      *model.User
	perawatan *model.Perawatan
	transaksi *model.Transaksi
)

func main() {
	config.Open()
	config.DB.AutoMigrate(perawatan, transaksi)
	config.DB.AutoMigrate(user, kambing)
	e := echo.New()

	var middlewareJWT = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/signin", controller.Login)
	e.POST("/signup", controller.CreateUser)
	e.GET("/kambing/:id", controller.GetAllKambing, middlewareJWT)
	e.POST("/kambing", controller.CreateKambingController, middlewareJWT)
	e.PUT("/kambing/:id", controller.UpdateKambingController, middlewareJWT)
	e.DELETE("/kambing/:id", controller.DeleteKambingController, middlewareJWT)
	e.GET("/perawatan/:id", controller.GetAllPerawatanFromKambing, middlewareJWT)
	e.POST("/perawatan", controller.CreatePerawatanFromKambing, middlewareJWT)
	e.PUT("/perawatan/:id", controller.UpdatePerawatanController, middlewareJWT)
	e.DELETE("/perawatan/:id", controller.DeletePerawatanController, middlewareJWT)
	e.GET("/transaksi/:id", controller.GetAllTransaksi, middlewareJWT)
	e.POST("/transaksi", controller.CreateTransaksi, middlewareJWT)
	e.PUT("/transaksi/:id", controller.UpdateTransaksiController, middlewareJWT)
	e.DELETE("/transaksi/:id", controller.DeleteTransaksiController, middlewareJWT)
	e.Logger.Fatal(e.Start(":8080"))
}
