package inventorycontroller

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var inventory []models.Inventory
	models.DB.Debug().Preload("Room").Preload("Employee").Preload("Category").Find(&inventory)
	models.DB.Preload("Room").Preload("Employee").Preload("Category").Find(&inventory)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"inventory": inventory})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	id := mux.Vars(r)["kode_aset"]

	models.DB.Debug().Preload("Room").Preload("Employee").Preload("Category").Find(&inventory, "kode_aset = ?", id)
	if err := models.DB.Preload("Room").Preload("Employee").Preload("Category").Where("kode_aset = ?", id).First(&inventory).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Aset tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"inventory": inventory})
}

// func Create(w http.ResponseWriter, r *http.Request) {
// 	var inventory models.Inventory

// 	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
// 		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Create a new inventory record
// 	models.DB.Create(&inventory)

// 	// Create a new history pemakaian record
// 	createHistoryPemakaian(inventory, "", inventory.EmployeeID, "", inventory.RoomID)

// 	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data Berhasil Dibuat"})
// }

func Create(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory

	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Calculate Depreciation
	depreciation := (inventory.Price - (inventory.Price / 4)) / inventory.UsefulLife

	// Calculate Yearly values
	inventory.Depreciation = depreciation
	inventory.Year1 = inventory.Price - depreciation
	inventory.Year2 = inventory.Year1 - depreciation
	inventory.Year3 = inventory.Year2 - depreciation
	inventory.Year4 = inventory.Year3 - depreciation

	// Create a new inventory record
	models.DB.Create(&inventory)

	// Create a new history pemakaian record
	createHistoryPemakaian(inventory, "", inventory.EmployeeID, "", inventory.RoomID)

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Aset Berhasil Dibuat"})
}



func Update(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	id := mux.Vars(r)["kode_aset"]

	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Get the old inventory record
	var oldInventory models.Inventory
	models.DB.First(&oldInventory, "kode_aset = ?", id)

	// Update the inventory record
	if models.DB.Model(&models.Inventory{}).Where("kode_aset = ?", id).Updates(&inventory).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Tidak dapat memperbarui Aset"})
		return
	}

	// Create a new history pemakaian record
	createHistoryPemakaian(oldInventory, oldInventory.EmployeeID, inventory.EmployeeID, oldInventory.RoomID, inventory.RoomID)

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data Berhasil Diperbaharui"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	id := mux.Vars(r)["kode_aset"]

	if err := models.DB.First(&inventory, "kode_aset = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Aset tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&inventory).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus Aset"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data berhasil dihapus"})
}

func createHistoryPemakaian(inventory models.Inventory, oldEmployeeID, newEmployeeID, oldRoomID, newRoomID string) {
	historyPemakaian := models.ReportHistoryPemakaian{
		OldEmployeeID: oldEmployeeID,
		NewEmployeeID: newEmployeeID,
		OldRoom:       oldRoomID,
		NewRoom:       newRoomID,
		UsageDate:     inventory.UpdatedAt,
		AssetCode:     inventory.AssetCode,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Create a new history pemakaian record
	models.DB.Create(&historyPemakaian)
}
