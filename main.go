package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/06202003/apiInventory/middlewares"

	"github.com/06202003/apiInventory/controllers/authcontroller"
	"github.com/06202003/apiInventory/controllers/categorycontroller"
	"github.com/06202003/apiInventory/controllers/employeecontroller"
	"github.com/06202003/apiInventory/controllers/inventorycontroller"
	"github.com/06202003/apiInventory/controllers/locationcontroller"
	"github.com/06202003/apiInventory/controllers/logkerusakancontroller"
	"github.com/06202003/apiInventory/controllers/portfoliocontroller"
	"github.com/06202003/apiInventory/controllers/portfolioemployeecontroller"
	"github.com/06202003/apiInventory/controllers/reporthistorykerusakancontroller"
	"github.com/06202003/apiInventory/controllers/reporthistorypemakaiancontroller"
	"github.com/06202003/apiInventory/controllers/reporthistoryperbaikancontroller"
	"github.com/06202003/apiInventory/controllers/roomcontroller"
	"github.com/06202003/apiInventory/controllers/skillcontroller"
	"github.com/06202003/apiInventory/controllers/skillemployeecontroller"
	"github.com/06202003/apiInventory/controllers/usagecontroller"
	"github.com/06202003/apiInventory/models"

	"github.com/gorilla/mux"
)

var version = "1.0.0"

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/categories", categorycontroller.Index).Methods("GET")
	api.HandleFunc("/categories/{id_kategori}", categorycontroller.Show).Methods("GET")
	api.HandleFunc("/categories", categorycontroller.Create).Methods("POST")
	api.HandleFunc("/categories/{id_kategori}", categorycontroller.Update).Methods("PUT")
	api.HandleFunc("/categories/{id_kategori}", categorycontroller.Delete).Methods("DELETE")

	api.HandleFunc("/employees", employeecontroller.Index).Methods("GET")
	api.HandleFunc("/employees/{nomor_induk}", employeecontroller.Show).Methods("GET")
	api.HandleFunc("/employees", employeecontroller.Create).Methods("POST")
	api.HandleFunc("/employees/{nomor_induk}", employeecontroller.Update).Methods("PUT")
	api.HandleFunc("/employees/{nomor_induk}", employeecontroller.Delete).Methods("DELETE")
	

	api.HandleFunc("/skills", skillcontroller.Index).Methods("GET")
	api.HandleFunc("/skills/{id_skill}", skillcontroller.Show).Methods("GET")
	api.HandleFunc("/skills", skillcontroller.Create).Methods("POST")
	api.HandleFunc("/skills/{id_skill}", skillcontroller.Update).Methods("PUT")
	api.HandleFunc("/skills/{id_skill}", skillcontroller.Delete).Methods("DELETE")

	api.HandleFunc("/employeeSkills", skillemployeecontroller.Index).Methods("GET")
	api.HandleFunc("/employeeSkills/{id_skill_employee}", skillemployeecontroller.Show).Methods("GET")
	api.HandleFunc("/employeeSkills", skillemployeecontroller.Create).Methods("POST")
	api.HandleFunc("/employeeSkills/{id_skill_employee}", skillemployeecontroller.Update).Methods("PUT")
	api.HandleFunc("/employeeSkills/{id_skill_employee}", skillemployeecontroller.Delete).Methods("DELETE")
	api.HandleFunc("/employeeSkills/ViewByEmployeeSkill/{id_skill_employee}", skillemployeecontroller.ViewByEmployeeSkill).Methods("GET")
	
	api.HandleFunc("/portfolio", portfoliocontroller.Index).Methods("GET")
	api.HandleFunc("/portfolio/{id_portfolio}", portfoliocontroller.Show).Methods("GET")
	api.HandleFunc("/portfolio", portfoliocontroller.Create).Methods("POST")
	api.HandleFunc("/portfolio/{id_portfolio}", portfoliocontroller.Update).Methods("PUT")
	api.HandleFunc("/portfolio/{id_portfolio}", portfoliocontroller.Delete).Methods("DELETE")

	api.HandleFunc("/employeePortfolio", portfolioemployeecontroller.Index).Methods("GET")
	api.HandleFunc("/employeePortfolio/{id_portfolio_employee}", portfolioemployeecontroller.Show).Methods("GET")
	api.HandleFunc("/employeePortfolio", portfolioemployeecontroller.Create).Methods("POST")
	api.HandleFunc("/employeePortfolio/{id_portfolio_employee}", portfolioemployeecontroller.Update).Methods("PUT")
	api.HandleFunc("/employeePortfolio/{id_portfolio_employee}", portfolioemployeecontroller.Delete).Methods("DELETE")
	api.HandleFunc("/employeePortfolio/ViewByEmployeePortfolio/{id_portfolio_employee}", portfolioemployeecontroller.ViewByEmployeePortfolio).Methods("GET")

	api.HandleFunc("/inventories", inventorycontroller.Index).Methods("GET")
	api.HandleFunc("/inventories/{kode_aset}", inventorycontroller.Show).Methods("GET")
	api.HandleFunc("/inventories", inventorycontroller.Create).Methods("POST")
	api.HandleFunc("/inventories/{kode_aset}", inventorycontroller.Update).Methods("PUT")
	api.HandleFunc("/inventories/{kode_aset}", inventorycontroller.Delete).Methods("DELETE")

	api.HandleFunc("/usages", usagecontroller.Index).Methods("GET")
	api.HandleFunc("/usages/{id_pemakaian}", usagecontroller.Show).Methods("GET")
	api.HandleFunc("/usages", usagecontroller.Create).Methods("POST")
	api.HandleFunc("/usages/{id_pemakaian}", usagecontroller.Update).Methods("PUT")
	api.HandleFunc("/usages/{id_pemakaian}", usagecontroller.Delete).Methods("DELETE")

	// New routes for viewing usages by room and by employee
	api.HandleFunc("/Usageroom/{id_ruangan}", usagecontroller.ViewByRoom).Methods("GET")
	api.HandleFunc("/Usageemployee/{nomor_induk}", usagecontroller.ViewByEmployee).Methods("GET")

	api.HandleFunc("/usageHistories", reporthistorypemakaiancontroller.Index).Methods("GET")
	api.HandleFunc("/usageHistories/{id}", reporthistorypemakaiancontroller.Show).Methods("GET")

	api.HandleFunc("/logProblem", logkerusakancontroller.Index).Methods("GET")
	api.HandleFunc("/logProblem/{id}", logkerusakancontroller.Show).Methods("GET")

	api.HandleFunc("/repairHistories", reporthistoryperbaikancontroller.Index).Methods("GET")
	api.HandleFunc("/repairHistories/{id_perbaikan}", reporthistoryperbaikancontroller.Show).Methods("GET")
	api.HandleFunc("/repairHistories", reporthistoryperbaikancontroller.Create).Methods("POST")
	api.HandleFunc("/repairHistories/{id_perbaikan}", reporthistoryperbaikancontroller.Update).Methods("PUT")
	api.HandleFunc("/repairHistories/{id_perbaikan}", reporthistoryperbaikancontroller.Delete).Methods("DELETE")

	api.HandleFunc("/problemHistories", reporthistorykerusakancontroller.Index).Methods("GET")
	api.HandleFunc("/problemHistories/{id}", reporthistorykerusakancontroller.Show).Methods("GET")
	api.HandleFunc("/problemHistories", reporthistorykerusakancontroller.Create).Methods("POST")
	api.HandleFunc("/problemHistories/{id}", reporthistorykerusakancontroller.Update).Methods("PUT")
	api.HandleFunc("/problemHistories/{id}", reporthistorykerusakancontroller.Delete).Methods("DELETE")

	api.HandleFunc("/rooms", roomcontroller.Index).Methods("GET")
	api.HandleFunc("/rooms/{id_ruangan}", roomcontroller.Show).Methods("GET")
	api.HandleFunc("/rooms", roomcontroller.Create).Methods("POST")
	api.HandleFunc("/rooms/{id_ruangan}", roomcontroller.Update).Methods("PUT")
	api.HandleFunc("/rooms/{id_ruangan}", roomcontroller.Delete).Methods("DELETE")

	api.HandleFunc("/locations", locationcontroller.Index).Methods("GET")
	api.HandleFunc("/locations/{id_lokasi}", locationcontroller.Show).Methods("GET")
	api.HandleFunc("/locations", locationcontroller.Create).Methods("POST")
	api.HandleFunc("/locations/{id_lokasi}", locationcontroller.Update).Methods("PUT")
	api.HandleFunc("/locations/{id_lokasi}", locationcontroller.Delete).Methods("DELETE")


	api.Use(middlewares.JWTMiddleware)

	printBanner()

	fmt.Printf("App Version %s\n", version)
	fmt.Printf("Happy Hacking 🧑‍💻")
	fmt.Printf("Happy Hacking 🧑‍💻")
	fmt.Printf(`
	
Created  by
Yezekiel David Setiawan

Supported by
Cristianto Tri Arthurito
Laurentius Gusti Ontoseno Panata Yudha
	`)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func printBanner() {
	banner := `
██╗    ██╗██╗████████╗                                                     
██║    ██║██║╚══██╔══╝                                                     
██║ █╗ ██║██║   ██║                                                        
██║███╗██║██║   ██║                                                        
╚███╔███╔╝██║   ██║                                                        
╚══╝╚══╝ ╚═╝   ╚═╝                                                        
																	
██╗███╗   ██╗██╗   ██╗███████╗███╗   ██╗████████╗ ██████╗ ██████╗ ██╗   ██╗
██║████╗  ██║██║   ██║██╔════╝████╗  ██║╚══██╔══╝██╔═══██╗██╔══██╗╚██╗ ██╔╝
██║██╔██╗ ██║██║   ██║█████╗  ██╔██╗ ██║   ██║   ██║   ██║██████╔╝ ╚████╔╝ 
██║██║╚██╗██║╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║   ██║   ██║██╔══██╗  ╚██╔╝  
██║██║ ╚████║ ╚████╔╝ ███████╗██║ ╚████║   ██║   ╚██████╔╝██║  ██║   ██║   
╚═╝╚═╝  ╚═══╝  ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝   
																	
███████╗██╗   ██╗███████╗████████╗███████╗███╗   ███╗                      
██╔════╝╚██╗ ██╔╝██╔════╝╚══██╔══╝██╔════╝████╗ ████║                      
███████╗ ╚████╔╝ ███████╗   ██║   █████╗  ██╔████╔██║                      
╚════██║  ╚██╔╝  ╚════██║   ██║   ██╔══╝  ██║╚██╔╝██║                      
███████║   ██║   ███████║   ██║   ███████╗██║ ╚═╝ ██║                      
╚══════╝   ╚═╝   ╚══════╝   ╚═╝   ╚══════╝╚═╝     ╚═╝                      
																										
	`
	fmt.Println(banner)
}
