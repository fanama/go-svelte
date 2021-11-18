package main

import (
	"net/http"

	"github.com/fanama/go-svelte/Api/router"
)

func main() {

	http.HandleFunc("/hello", router.Hello)
	http.HandleFunc("/headers", router.Headers)
	http.HandleFunc("/get", router.GetData)       // ?db=test
	http.HandleFunc("/insert", router.InsertData) //?db=test

	http.ListenAndServe(":8090", nil)
}
