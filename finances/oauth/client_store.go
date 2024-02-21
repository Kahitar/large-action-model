package oauth

import (
	"database/sql"
	"finances/database"
	"fmt"
	"log"

	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

const (
    OAUTH_CLIENTS_DB = "oauth-clients"
    OAUTH_CLIENTS_TABLE = "clients"
)

type TursoClientStore struct {
    db *sql.DB
}

func NewTursoClientStore() (*TursoClientStore, error) {
    dbInfo, err := getOauthDb()
    if err != nil {
        return nil, err
    }
    dbParams := database.NewDBParams(dbInfo)
    db := database.CreateDbConnection(dbParams)
    return &TursoClientStore{db: db}, nil
}

func (s *TursoClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
    var client models.Client
    err := s.db.QueryRow("SELECT id, secret, domain FROM clients WHERE id = ?", id).Scan(&client.ID, &client.Secret, &client.Domain)
    if err != nil {
        return nil, err
    }
    return &client, nil
}

func (s *TursoClientStore) Set(id string, client oauth2.ClientInfo) error {
    _, err := s.db.Exec("INSERT INTO clients (id, secret, domain) VALUES (?, ?, ?)", id, client.GetSecret(), client.GetDomain())
    return err
}

func getOauthDb() (database.DatabaseInfo, error) {
    databaseParams := database.NewDatabaseParams(OAUTH_CLIENTS_DB)
    if !database.DatabaseExists(databaseParams) {
        _, err := database.CreateDB(databaseParams)
        if err != nil {
            return database.DatabaseInfo{}, err
        }
    }
    dbInfo, err := database.GetDbInfo(databaseParams)
    if err != nil {
        return database.DatabaseInfo{}, err
    }
    if err := createClientTable(dbInfo); err != nil {
        log.Printf("error creating client table: %v", err)
        return database.DatabaseInfo{}, err
    }
    return dbInfo, nil
}

func createClientTable(dbInfo database.DatabaseInfo) error {
    dbParams := database.NewDBParams(dbInfo)
    db := database.CreateDbConnection(dbParams)
    defer database.CloseDbConnection(db)

    log.Printf("creating table %s", OAUTH_CLIENTS_TABLE)
    sqlStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id TEXT PRIMARY KEY, secret TEXT, domain TEXT)", OAUTH_CLIENTS_TABLE)
    _, err := db.Exec(sqlStmt)
    return err
}

