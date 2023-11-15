package main

import (
	"log"
	"net/http"

	"github.com/06202003/apiInventory/middlewares"

	"github.com/06202003/apiInventory/controllers/authcontroller"
	"github.com/06202003/apiInventory/controllers/categorycontroller"
	"github.com/06202003/apiInventory/controllers/employeecontroller"
	"github.com/06202003/apiInventory/controllers/roomcontroller"
	"github.com/06202003/apiInventory/controllers/inventorycontroller"
	"github.com/06202003/apiInventory/controllers/reporthistorypemakaiancontroller"
	"github.com/06202003/apiInventory/controllers/reporthistoryperbaikancontroller"
	"github.com/06202003/apiInventory/controllers/usagecontroller"
	"github.com/06202003/apiInventory/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	
	api.HandleFunc("/categories", categorycontroller.Index).Methods("GET")
	api.HandleFunc("/categories/show/{id_kategori}", categorycontroller.Show).Methods("GET")
	api.HandleFunc("/categories", categorycontroller.Create).Methods("POST")
	api.HandleFunc("/categories/update/{id_kategori}", categorycontroller.Update).Methods("PUT")
	api.HandleFunc("/categories/delete/{id_kategori}", categorycontroller.Delete).Methods("DELETE")

	api.HandleFunc("/employees", employeecontroller.Index).Methods("GET")
	api.HandleFunc("/employees/show/{nomor_induk}", employeecontroller.Show).Methods("GET")
	api.HandleFunc("/employees", employeecontroller.Create).Methods("POST")
	api.HandleFunc("/employees/update/{nomor_induk}", employeecontroller.Update).Methods("PUT")
	api.HandleFunc("/employees/delete/{nomor_induk}", employeecontroller.Delete).Methods("DELETE")

	api.HandleFunc("/inventory", inventorycontroller.Index).Methods("GET")
	api.HandleFunc("/inventory/show/{kode_aset}", inventorycontroller.Show).Methods("GET")
	api.HandleFunc("/inventory", inventorycontroller.Create).Methods("POST")
	api.HandleFunc("/inventory/update/{kode_aset}", inventorycontroller.Update).Methods("PUT")
	api.HandleFunc("/inventory/delete/{kode_aset}", inventorycontroller.Delete).Methods("DELETE")

	api.HandleFunc("/historyPemakaian", reporthistorypemakaiancontroller.Index).Methods("GET")
	api.HandleFunc("/historyPemakaian/show/{id}", reporthistorypemakaiancontroller.Show).Methods("GET")

	api.HandleFunc("/historyPerbaikan", reporthistoryperbaikancontroller.Index).Methods("GET")
	api.HandleFunc("/historyPerbaikan/show/{id}", reporthistoryperbaikancontroller.Show).Methods("GET")
	api.HandleFunc("/historyPerbaikan", reporthistoryperbaikancontroller.Create).Methods("POST")
	api.HandleFunc("/historyPerbaikan/update/{id}", reporthistoryperbaikancontroller.Update).Methods("PUT")
	api.HandleFunc("/historyPerbaikan/delete/{id}", reporthistoryperbaikancontroller.Delete).Methods("DELETE")

	api.HandleFunc("/room", roomcontroller.Index).Methods("GET")
	api.HandleFunc("/room/show/{id_ruangan}", roomcontroller.Show).Methods("GET")
	api.HandleFunc("/room", roomcontroller.Create).Methods("POST")
	api.HandleFunc("/room/update/{id_ruangan}", roomcontroller.Update).Methods("PUT")
	api.HandleFunc("/room/delete/{id_ruangan}", roomcontroller.Delete).Methods("DELETE")

	api.HandleFunc("/usage", usagecontroller.Index).Methods("GET")
	api.HandleFunc("/usage/show/{id_pemakaian}", usagecontroller.Show).Methods("GET")
	api.HandleFunc("/usage", usagecontroller.Create).Methods("POST")
	api.HandleFunc("/usage/update/{id_pemakaian}", usagecontroller.Update).Methods("PUT")
	api.HandleFunc("/usage/delete/{id_pemakaian}", usagecontroller.Delete).Methods("DELETE")


	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
