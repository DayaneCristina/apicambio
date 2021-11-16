package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Create Read Update Delete (CRUD)

	fmt.Println("API Cambio")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))

}
