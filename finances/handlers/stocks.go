package handlers

import (
	"finances/database"
	"net/http"
)

const (
    DB_NAME = "first-test-db"
)

func AddStockHandler(w http.ResponseWriter, r *http.Request) {
    // do something
    dbParams := database.NewDBParams(DB_NAME)
    db := database.CreateDbConnection(dbParams)
    defer database.CloseDbConnection(db)

    existParams := database.NewDatabaseParams(DB_NAME)
    print("Exists: %v\n", database.DatabaseExists(existParams))
    
    // database.QueryUsers(db) 

}
