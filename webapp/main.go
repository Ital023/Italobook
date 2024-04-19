package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Rodando web app")


	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}