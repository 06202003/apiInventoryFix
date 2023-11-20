package models

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Open a database connection
    database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/inventorysuperfix"))
    if err != nil {
        panic(err)
    }

    // Perform table migrations
    autoMigrateAllTables(database)

    // Assign the database connection to the global variable
    DB = database
}

func autoMigrateAllTables(db *gorm.DB) {
    // Define all the models
    models := []interface{}{
        &User{},
        &Employee{},
        &Category{},
        &Inventory{},
        &ReportHistoryPemakaian{},
        &ReportHistoryPerbaikan{},
        &Room{},
    }

    // AutoMigrate all tables
    for _, model := range models {
        if err := db.AutoMigrate(model); err != nil {
            panic(err)
        }
    }
}
