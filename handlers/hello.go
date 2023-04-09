package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	resp := make(map[string]string)
	resp["message"] = "Hello, HTTP"
	jsonResp, err := json.Marshal(resp)
	if (err != nil) {
		fmt.Printf("Error occured in JSON marshal. Err: %s\n", err)
	}
	w.Write(jsonResp)
}