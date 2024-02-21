package handlers

import (
	"finances/database"
	"log"
	"net/http"
)

const (
    DB_NAME = "userxyz"
)

func AddStockHandler(w http.ResponseWriter, r *http.Request) {
    dbParams := database.NewDBParams(DB_NAME)
    db := database.CreateDbConnection(dbParams)
    defer database.CloseDbConnection(db)

    databaseParams := database.NewDatabaseParams(DB_NAME)
    if !database.DatabaseExists(databaseParams) {
        if ok := database.CreateDB(databaseParams); !ok {
            http.Error(w, "Failed to create database", http.StatusInternalServerError)
            return
        }
    }

    log.Println("All good!")
}
