package database

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type databaseParams struct {
    token string
    organizationName string
    databaseName string
    group string
}

func NewDatabaseParams(databaseName string) databaseParams {
    return databaseParams{
        token: os.Getenv("DB_TOKEN"),
        organizationName: os.Getenv("ORGANIZATION"),
        databaseName: databaseName,
        group: os.Getenv("GROUP"),
    }
}

func DatabaseExists(params databaseParams) bool {
    requestUrl := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s", params.organizationName, params.databaseName)
    req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
    if err != nil {
        log.Printf("error in create retrieve database request: %v", err)
        return false
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", params.token))

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("error in retrieve database request: %v", err)
        return false
    }
    
    // resBody, err := io.ReadAll(res.Body)
    // fmt.Printf("body: %s\n", resBody)

    return res.StatusCode == 200
}


func CreateDB(params databaseParams) bool {
    // curl -L -X POST 'https://api.turso.tech/v1/organizations/{organizationName}/databases' \
    // -H 'Authorization: Bearer TOKEN' \
    // -H 'Content-Type: application/json' \
    // -d '{
    //   "name": "new-database",
    //   "group": "default"
    // }'

    requestUrl := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", params.organizationName)
    req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
    if err != nil {
        log.Printf("error in create retrieve database request: %v", err)
        return false
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", params.token))

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("error in retrieve database request: %v", err)
        return false
    }
    
    // resBody, err := io.ReadAll(res.Body)
    // fmt.Printf("body: %s\n", resBody)

    return res.StatusCode == 200
}

