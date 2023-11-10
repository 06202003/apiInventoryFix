package usagecontroller

import (
	"encoding/json"
	"net/http"
	"github.com/06202003/apiInventory/helper"
	"github.com/06202003/apiInventory/models"
	"github.com/gorilla/mux"
)

// Index handles the request to retrieve all usages
func Index(w http.ResponseWriter, r *http.Request) {
	var usages []models.Usage
	models.DB.Debug().Preload("Inventory").Preload("Inventory.Category").Preload("Inventory.Employee").Preload("Room").Find(&usages)
	models.DB.Preload("Inventory").Preload("Inventory.Category").Preload("Inventory.Employee").Preload("Room").Find(&usages)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usages": usages})
}

// Show handles the request to retrieve a specific usage by ID
func Show(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage
	id := mux.Vars(r)["id_pemakaian"]

	if err := models.DB.Preload("Inventory").Preload("Inventory.Category").Preload("Inventory.Employee").Preload("Room").First(&usage, id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Usage not found"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usage": usage})
}

// Create handles the request to create a new usage
func Create(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage

	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	models.DB.Create(&usage)
	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"usage": usage})
}

// Update handles the request to update an existing usage by ID
func Update(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage
	id := mux.Vars(r)["id_pemakaian"]

	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Enable SQL debugging
	models.DB = models.DB.Debug()


	// Perform the update
	if models.DB.Model(&usage).Where("id_pemakaian = ?", id).Updates(&usage).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Failed to update usage"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data updated successfully"})
}



// Delete handles the request to delete a specific usage by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	var usage models.Usage
	id := mux.Vars(r)["id_pemakaian"]

	// Check if the usage exists
	if err := models.DB.First(&usage, id).Error; err != nil {
		helper.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "Usage not found"})
		return
	}

	// Delete the usage with a specific condition
	if err := models.DB.Where("id_pemakaian = ?", id).Delete(&usage).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"message": "Failed to delete usage"})
		return
	}

	helper.ResponseJSON(w, http.StatusOK, map[string]interface{}{"message": "Data deleted successfully"})
}

