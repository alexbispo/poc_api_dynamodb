package main

import (
	"encoding/json"
	"net/http"
	// "strconv"
	"path"
)

type Card struct {
	indice_pk string `json:"indice_pk"`
	indice_sk string `json:"indice_sk"`
	account_id int `json:"account_id"`
	card_id int `json:"card_id"`
	contract_id int `json:"contract_id"`
	main_account_id int `json:"main_account_id"`
	external_code string `json:"external_code"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/cards/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err = handleGet(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve a card
// GET /accounts/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// id, err := strconv.Atoi(path.Base(r.URL.Path))
	id := path.Base(r.URL.Path)

	// if err != nil {
	// 	return
	// }
	card, err := RetrieveCardById(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&card, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
