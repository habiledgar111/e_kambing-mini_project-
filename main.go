package main

import (
	"mini_project/config"
	"mini_project/controller"
	"mini_project/model"
	"net/http"

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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/signin", controller.Login)
	e.POST("/signup", controller.CreateUser)
	e.GET("/kambing/:id", controller.GetAllKambing)
	e.POST("/kambing", controller.CreateKambingController)
	e.PUT("/kambing/:id", controller.UpdateKambingController)
	e.GET("/perawatan/:id", controller.GetAllPerawatanFromKambing)
	e.POST("/perawatan", controller.CreatePerawatanFromKambing)
	e.PUT("/perawatan/:id", controller.UpdatePerawatanController)
	e.DELETE("/perawatan/:id", controller.DeletePerawatanController)
	e.GET("/transaksi/:id", controller.GetAllTransaksi)
	e.POST("/transaksi", controller.CreateTransaksi)
	e.Logger.Fatal(e.Start(":1323"))
}
