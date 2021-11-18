package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanama/go-svelte/Api/database"
)

func InsertData(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["db"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}

	db := string(keys[0])

	// Declare a new Person struct.
	var p map[string]string

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...

	err = database.Write(db, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)

		return
	}

	fmt.Fprintf(w, "insertin Data : %+v into %v success!!!", p, db)
}

func GetData(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["db"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Querry 'key' is missing")
		return
	}

	db := string(keys[0])

	results, err := database.Read(db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)

		return
	}

	jsonStr, err := json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)

		return
	}

	fmt.Fprintf(w, string(jsonStr))

	return
}
