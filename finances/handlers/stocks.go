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
    platformParams := database.NewPlatformParamsFromEnv()
    dbInfo, err := database.GetDbInfo(platformParams, DB_NAME)
    if err != nil {
        log.Println("Error getting database info:", err)
        http.Error(w, "Failed to get database info", http.StatusInternalServerError)
        return
    }

    dbParams := database.NewDBParams(dbInfo)
    db := database.CreateDbConnection(dbParams)
    defer database.CloseDbConnection(db)

    databaseParams := database.NewPlatformParamsFromEnv()
    if !database.DatabaseExists(databaseParams, DB_NAME) {
        if _, err := database.CreateDB(databaseParams, DB_NAME); err != nil {
            http.Error(w, "Failed to create database", http.StatusInternalServerError)
            return
        }
    }

    log.Println("All good!")
}
