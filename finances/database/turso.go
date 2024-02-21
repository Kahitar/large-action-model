package database

import ()

type DatabaseInfo struct {
    Token string
    Name string
}

type ErrorResponse struct {
    Code string `json:"code"`
    Error string `json:"error"`
}

type TokenResponse struct {
    Jwt string `json:"jwt"`
}

