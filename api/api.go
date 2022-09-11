package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nguyenvulong/go-bank-backend/helpers"
	"github.com/nguyenvulong/go-bank-backend/users"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {

	}

	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	log.Info().Msg("App is working on port :8888")
	log.Fatal().Err(http.ListenAndServe(":8888", router)).
	Msg("Bank Web App closed")

}
