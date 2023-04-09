package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/prestonchoate/go-user-service/models"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]string)

	if (r.Method != "POST") {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp["error"] = "Not Allowed"
		jsonResp, _ := generateResponse(resp)
		w.Write(jsonResp)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestData models.LoginRequest
	err := decoder.Decode(&requestData)

	if (err != nil) {
		fmt.Println("failed to decode login request")
		w.WriteHeader(http.StatusBadRequest)
		resp["error"] = "Bad Request"
		jsonResp, _ := generateResponse(resp)
		w.Write(jsonResp)
		return
	}

	user, err := models.CheckLogin(requestData.Username, requestData.Password)
	if (err != nil) {
		resp["error"] = err.Error()
		jsonResp, _ := generateResponse(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}
	session, err := models.GetSession(user)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		resp["error"] = err.Error()
		jsonResp, _ := generateResponse(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp["data"] = session.AuthKey
	jsonResp, _ := generateResponse(resp)
	w.Write(jsonResp)
}

func generateResponse(resp map[string]string) ([]byte, error){
	jsonResp, err := json.Marshal(resp)
	if (err != nil) {
		fmt.Printf("Error occured in JSON marshal. Err: %s\n", err)
		return nil, errors.New("could not marshal JSON response")
	}
	return jsonResp, nil
}