package main

import (
	"html/template"
	"log"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/index.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	Produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, Bem Bonita", Preco: 39, Quantidade: 2},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito Bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", Produtos)
}
