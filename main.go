package main

import (
	"fmt"
	"github.com/ojaoferreira/go-rest-api/database"
	"github.com/ojaoferreira/go-rest-api/src/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando uma aplicação API com GO")
	r := routes.HandleRequest()
	database.ConectaComBancoDeDados()
	log.Fatal(http.ListenAndServe(":8000", r))
}
