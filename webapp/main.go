package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/utils"
)

func main() {
	
	utils.CarregarTemplates()

	r := router.Gerar()

	fmt.Println("Rodando web app")
	log.Fatal(http.ListenAndServe(":3000", r))
}