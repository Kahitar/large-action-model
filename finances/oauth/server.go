package oauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func InitializeOauthServer(mux *http.ServeMux) {
    manager := manage.NewDefaultManager()
    manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

    // token memory store - TODO: replace with turso token store
    manager.MustTokenStorage(store.NewMemoryTokenStore())

    // client memory store
    // clientStore := store.NewClientStore()
    clientStore, err := NewTursoClientStore()
    if err != nil {
        log.Fatal(err)
    }

    manager.MapClientStorage(clientStore)

    srv := server.NewDefaultServer(manager)
    srv.SetAllowGetAccessRequest(true)
    srv.SetClientInfoHandler(server.ClientFormHandler)
    manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

    srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
        log.Println("Internal Error:", err.Error())
        return
    })

    srv.SetResponseErrorHandler(func(re *errors.Response) {
        log.Println("Response Error:", re.Error.Error())
    })

    mux.HandleFunc("GET /token", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Token Request: %v", r)
        srv.HandleTokenRequest(w, r)
    })

    mux.HandleFunc("GET /credentials", func(w http.ResponseWriter, r *http.Request) {
        clientId := uuid.New().String()[:8]
        clientSecret := uuid.New().String()[:8]
        err := clientStore.Set(clientId, &models.Client{
            ID:     clientId,
            Secret: clientSecret,
            Domain: "http://localhost:9096",
        })
        if err != nil {
            fmt.Println(err.Error())
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
    })

    mux.HandleFunc("GET /protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, I'm protected\n"))
    }, srv))
}

func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, err := srv.ValidationBearerToken(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        f.ServeHTTP(w, r)
    })
}
