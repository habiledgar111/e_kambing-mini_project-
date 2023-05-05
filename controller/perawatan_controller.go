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
// 	DB_perawatan *gorm.DB
// 	perawatan    *model.Perawatan
// )

// func init() {
// 	dsn := "root:Mbahbambang123@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	DB_perawatan = db
// 	DB_perawatan.AutoMigrate(perawatan)
// }

func GetAllPerawatanFromKambing(c echo.Context) error {
	id_kambing := c.Param("id")
	id_kambing_int, _ := strconv.Atoi(id_kambing)
	perawatans, err := model.GetAllPerawatan(id_kambing_int)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if len(perawatans.Perawatans) <= 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "dont have perawatan",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all perawatan",
		"kambing": perawatans,
	})
}

func CreatePerawatanFromKambing(c echo.Context) error {
	var perawatans model.Perawatan
	// id_kambing,_  := strconv.Atoi(c.Param("id"))
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	harga, _ := strconv.ParseFloat(fmt.Sprintf("%v", json_map["harga"]), 3)
	kambingID, _ := strconv.Atoi(fmt.Sprintf("%v", json_map["KambingID"]))
	perawatans = model.Perawatan{
		Name:       fmt.Sprintf("%v", json_map["name"]),
		Keterangan: fmt.Sprintf("%v", json_map["keterangan"]),
		KambingID:  uint(kambingID),
		Harga:      harga,
		Tanggal:    time.Now(),
	}
	result := model.CreatePerawatan(perawatans)

	transaksi := model.Transaksi{
		Name:       ("Perawatan" + fmt.Sprintf("%v", json_map["KambingID"])),
		Keterangan: perawatans.Keterangan,
		KambingID:  perawatans.KambingID,
		Tanggal:    perawatans.Tanggal,
	}
	result_transaksi := model.CreateTransaksifromPerawatan(transaksi)

	if result_transaksi <= 0 {
		return c.JSON(http.StatusInternalServerError, "error save transaksi")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": result,
	})
}
