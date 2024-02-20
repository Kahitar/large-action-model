package handlers

import (
	"finances/database"
	"net/http"
)


func AddStockHandler(w http.ResponseWriter, r *http.Request) {
    // do something
    dbParams := database.NewDBParams("first-test-db")
    db := database.CreateDbConnection(dbParams)
    defer database.CloseDbConnection(db)
    database.QueryUsers(db) 
}
