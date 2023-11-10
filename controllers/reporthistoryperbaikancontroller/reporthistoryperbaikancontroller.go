package reporthistoryperbaikancontroller

import (
	"encoding/json"
	"net/http"

	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var reportHistories []models.ReportHistoryPerbaikan
	if err := models.DB.Debug().Preload("Inventory").Preload("Inventory.Category").Preload("Inventory.Employee").Find(&reportHistories).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menarik data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"ReportHistories": reportHistories})
}

func Show(w http.ResponseWriter, r *http.Request) {
	var reportHistory models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := models.DB.Preload("Inventory").Preload("Inventory.Category").Preload("Inventory.Employee").First(&reportHistory, id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data tidak ditemukan"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"ReportHistory": reportHistory})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var reportHistoryPerbaikan models.ReportHistoryPerbaikan

	if err := json.NewDecoder(r.Body).Decode(&reportHistoryPerbaikan); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal membuat data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"ReportHistoryPerbaikan": reportHistoryPerbaikan})
}

func Update(w http.ResponseWriter, r *http.Request) {
	var reportHistoryPerbaikan models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&reportHistoryPerbaikan); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&reportHistoryPerbaikan).Where("id = ?", id).Updates(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Gagal mengupdate data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data berhasil diupdate"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var reportHistoryPerbaikan models.ReportHistoryPerbaikan
	id := mux.Vars(r)["id"]

	if err := models.DB.First(&reportHistoryPerbaikan, "id = ?", id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&reportHistoryPerbaikan).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus data"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data berhasil dihapus"})
}

