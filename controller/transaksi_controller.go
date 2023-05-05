package controller

import (
	"encoding/json"
	"fmt"
	"mini_project/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateTransaksi(c echo.Context) error {
	var transaksi model.Transaksi
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	userID, err := strconv.Atoi(fmt.Sprintf("%v", json_map["userID"]))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "userID cant string")
	}
	transaksi = model.Transaksi{
		Name:       fmt.Sprintf("%v", json_map["name"]),
		Keterangan: fmt.Sprintf("%v", json_map["keterangan"]),
		UserID:     userID,
		Tanggal:    time.Now(),
	}

	result := model.CreateTransaksifromUser(transaksi)
	if result <= 0 {
		return c.JSON(http.StatusInternalServerError, "server error cant save data")
	}

	return c.JSON(http.StatusOK, "success save data")
}

func GetAllTransaksi(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "userID cant string")
	}
	user, err := model.GetAllTransaksifromUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "cant get data")
	}
	if len(user.Transaksis) <= 0 {
		return c.JSON(http.StatusOK, "user dont have transaksi")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data",
		"user":    "trasaksi",
	})
}
