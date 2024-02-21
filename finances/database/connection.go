package database 

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
    DB_URL = "libsql://[DATABASE]-kahitar.turso.io?authToken=[TOKEN]"
)

type DBParams struct {
    Url string
}

func NewDBParams(dbInfo DatabaseInfo) DBParams {
    url := strings.Replace(DB_URL, "[DATABASE]", dbInfo.Name, 1)
    url = strings.Replace(url, "[TOKEN]", dbInfo.Token, 1)
    return DBParams{
        Url: url,
    }
}

func CreateDbConnection(params DBParams) *sql.DB {
    db, err := sql.Open("libsql", params.Url)
    if err != nil {
        log.Fatalf("failed to open db %s: %s\n", params.Url, err)
    }
    log.Printf("created db connection")
    return db
}

func CloseDbConnection(db *sql.DB) {
    db.Close()
    log.Println("closed db connection")
}
