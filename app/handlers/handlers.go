package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/yousysadmin/nsq-auth/app"
	"github.com/yousysadmin/nsq-auth/app/helpers"
	"github.com/yousysadmin/nsq-auth/app/models"
	"log"
	"net/http"
)

// PingHandler healthcheck endpoint handler
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
	return
}

// AuthHandler auth endpoint handler
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	remoteClientAddr := r.FormValue("remote_ip")
	tls := r.FormValue("tls")
	secret := r.FormValue("secret")

	identity, err := app.AppConfig.FindIdentityBySecret(secret)
	if err != nil {
		log.Printf("ERROR: unsuccefully auth: %s", err)
		helpers.HttpError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	state := models.StateResponse{
		TTL:            3600,
		Authorizations: identity.Authorizations,
		Identity:       identity.Identity,
		IdentityURL:    fmt.Sprintf("http://%s", app.AppConfig.BindAddr),
	}

	if err := json.NewEncoder(w).Encode(state); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Printf("INFO: succefully auth by identity: %s, remote ip: %s, tls: %v,  authorizations: %v", identity.Identity, remoteClientAddr, tls, state.Authorizations)
}
