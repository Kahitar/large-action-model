package database 

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
    DB_URL = "libsql://[DATABASE]-kahitar.turso.io?authToken=[TOKEN]"
)

type DBParams struct {
    Url string
}

func NewDBParams(dbName string) DBParams {
    token := os.Getenv("DB_TOKEN")
    url := strings.Replace(DB_URL, "[DATABASE]", dbName, 1)
    url = strings.Replace(url, "[TOKEN]", token, 1)
    return DBParams{
        Url: url,
    }
}

func CreateDbConnection(params DBParams) *sql.DB {
    fmt.Printf("%s", params.Url)
    db, err := sql.Open("libsql", params.Url)
    if err != nil {
        log.Fatalf("failed to open db %s: %s\n", params.Url, err)
    }
    log.Printf("Created db connection to %s\n", params.Url)
    return db
}

func CloseDbConnection(db *sql.DB) {
    db.Close()
    log.Println("Closed db connection")
}
