package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prestonchoate/go-user-service/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	jsonResp, err := json.Marshal(users)
	if (err != nil) {
		fmt.Printf("Error occured in JSON marshal. Err: %s\n", err)
	}
	w.Write(jsonResp)
}