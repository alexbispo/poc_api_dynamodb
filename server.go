package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"path"
)

type Account struct {
	AccountId int `json:"accountId"`
	ContractId int `json:"contractId"`
	MainAccountId int `json:"mainAccountId"`
	PersonId int `json:"personId"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/accounts/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err = handleGet(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve a account
// GET /accounts/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	account, err := RetrieveAccountById(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&account, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
