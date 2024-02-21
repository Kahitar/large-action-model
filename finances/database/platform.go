package database

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
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
        log.Printf("error in create retrieve database request: %v\n", err)
        return false
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", params.token))

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("error in retrieve database request: %v\n", err)
        return false
    }

    return res.StatusCode == 200
}


type databaseCreateRequest struct {
    Name string `json:"name"`
    Group string `json:"group"`
}

func CreateDB(params databaseParams) bool {
    fmt.Printf("Creating database: %s\n", params.databaseName)
    requestUrl := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", params.organizationName)
    createRequest := databaseCreateRequest{Name: params.databaseName, Group: "default"}
    data, err := json.Marshal(createRequest)
    if err != nil {
        log.Printf("error in marschal create database request: %v\n", err)
    }
    req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(data))
    if err != nil {
        log.Printf("error in database request: %v\n", err)
        return false
    }

    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", params.token))
    req.Header.Add("Content-Type", "application/json")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("error in create database request: %v\n", err)
        return false
    }

    resBody, err := io.ReadAll(res.Body)
    defer res.Body.Close()
    if err != nil {
        log.Printf("error sending create database request: %v\n", err)
    }
    fmt.Printf("body: %s\n", resBody)

    if res.StatusCode > 299 {
        log.Printf("unexpected response code from create request: %d\n", res.StatusCode)
        errResp := ErrorResponse{}
        json.Unmarshal(resBody, &errResp)
        log.Printf("error response: http code %d, message: %v\n", res.StatusCode, errResp.Error)
        return false
    }

    return res.StatusCode == 200
}

