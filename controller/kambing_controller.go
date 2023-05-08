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

// var (
// 	DB_kambing *gorm.DB
// 	kambing    *model.Kambing
// )

// func init() {
// 	//connect db
// 	dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	DB_kambing = db
// 	DB_kambing.AutoMigrate(kambing)
// }

func CreateKambingController(c echo.Context) error {
	var kambings model.Kambing
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	harga, _ := strconv.ParseFloat(fmt.Sprintf("%v", json_map["harga"]), 3)
	UserID, _ := strconv.Atoi(fmt.Sprintf("%v", json_map["UserID"]))
	kambings = model.Kambing{
		Name:        fmt.Sprintf("%v", json_map["name"]),
		TanggalBeli: time.Now(),
		Status:      "di kandang",
		Harga:       harga,
		UserID:      uint(UserID),
	}

	result, kambingid := model.CreateKambingModel(kambings)

	kambingID := strconv.Itoa(kambingid)
	transaksi := model.Transaksi{
		Name:       "membeli kambing",
		Keterangan: ("memebeli kambing - " + kambingID),
		KambingID:  uint(kambingid),
		Tanggal:    kambings.TanggalBeli,
	}
	result_kambing := model.CreateTransaksifromKambing(transaksi)

	if result_kambing <= 0 {
		return c.JSON(http.StatusInternalServerError, "cant save data transaksi")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": result,
	})
}

func GetAllKambing(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	UserandKambings, err := model.GetAllKambingsfromUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success get all data",
		"UserandKambings": UserandKambings,
	})
}

func UpdateKambingController(c echo.Context) error {
	id, err2 := strconv.Atoi(c.Param("id"))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, "cant refactor id")
	}
	var kambing model.Kambing
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	harga, _ := strconv.ParseFloat(fmt.Sprintf("%v", json_map["harga"]), 3)
	kambing = model.Kambing{
		ID:     id,
		Name:   fmt.Sprintf("%v", json_map["name"]),
		Status: fmt.Sprintf("%v", json_map["status"]),
		Harga:  harga,
	}

	result, update := model.UpdateKambing(id, kambing)

	if result <= 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":       "cant update data",
			"rows affected": result,
			"update":        update,
		})
	}

	transaksi := model.Transaksi{
		Name:      update.Name,
		Harga:     update.Harga,
		KambingID: uint(kambing.ID),
	}
	result_transaksi, update_transaksi := model.UpdateTransaksifromKambing(int(update.ID), transaksi)

	if result_transaksi <= 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":       " cant update transaksi",
			"rows affected": result,
			"update":        update_transaksi,
		})
	}

	return c.JSON(http.StatusOK, "sucess update data")
}

func DeleteKambingController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := model.DeleteKambing(id)

	if result <= 0 {
		return c.JSON(http.StatusInternalServerError, "cant delete data")
	}

	result_delete_transaksi := model.DeleteAllTransaksifromKambing(id)
	if result_delete_transaksi <= 0 {
		return c.JSON(http.StatusInternalServerError, "cant delete data")
	}

	return c.JSON(http.StatusOK, "success deelete data")
}
