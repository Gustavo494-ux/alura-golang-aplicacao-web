package main

import (
	"alura-golang-aplicacao-web/routes"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRotas()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
