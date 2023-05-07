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
	result, perawatanid := model.CreatePerawatan(perawatans)

	transaksi := model.Transaksi{
		Name:        ("Perawatan kambing - " + fmt.Sprintf("%v", json_map["KambingID"])),
		Keterangan:  perawatans.Keterangan,
		PerawatanID: perawatanid,
		Tanggal:     perawatans.Tanggal,
	}
	result_transaksi := model.CreateTransaksifromPerawatan(transaksi)

	if result_transaksi <= 0 {
		return c.JSON(http.StatusInternalServerError, "error save transaksi")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": result,
	})
}

func DeletePerawatanController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := model.DeletePerawatan(id)

	if result <= 0 {
		return c.JSON(http.StatusInternalServerError, "cant delete data")
	}

	result_delete_transaksi := model.DeleteAllTransaksifromPerawatan(id)

	if result_delete_transaksi <= 0 {
		return c.JSON(http.StatusInternalServerError, "cant delete data")
	}

	return c.JSON(http.StatusOK, "success deelete data")
}

func UpdatePerawatanController(c echo.Context) error {
	id, err2 := strconv.Atoi(c.Param("id"))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, "cant refactor id")
	}
	var perawatans model.Perawatan
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	harga, _ := strconv.ParseFloat(fmt.Sprintf("%v", json_map["harga"]), 3)
	perawatans = model.Perawatan{
		ID:         id,
		Name:       fmt.Sprintf("%v", json_map["name"]),
		Keterangan: fmt.Sprintf("%v", json_map["keterangan"]),
		Harga:      harga,
	}

	result, update := model.UpdatePerawatan(id, perawatans)

	if result <= 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":       "cant update data",
			"rows affected": result,
			"update":        update,
		})
	}

	transaksi := model.Transaksi{
		Name:        update.Name,
		Keterangan:  update.Keterangan,
		Harga:       update.Harga,
		PerawatanID: perawatans.ID,
	}
	result_transaksi, update_transaksi := model.UpdateTransaksifromPerawatan(int(update.ID), transaksi)

	if result_transaksi <= 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":       " cant update transaksi",
			"rows affected": result,
			"update":        update_transaksi,
		})
	}

	return c.JSON(http.StatusOK, "sucess update data")
}
