package main

import (
    "encoding/json"
    "math/rand"
    "net/http"

    _ "github.com/swaggo/swag/example/basic/docs" // docs is generated by Swag CLI, you need to import it.
)

// @title Hello World API
// @description This is a sample server for a hello world API.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}

// helloHandler returns a random greeting message.
// @Summary Return a random greeting
// @Description get a random greeting message
// @Tags hello
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [post]
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
        return
    }

    greetings := []string{"Hello, world!", "Hi there!", "Greetings, traveler!", "Howdy, partner!"}
    randomIndex := rand.Intn(len(greetings))

    response := map[string]string{"message": greetings[randomIndex]}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

