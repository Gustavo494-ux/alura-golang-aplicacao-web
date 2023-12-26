package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=Senha123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/index.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()

	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()
	defer db.Close()

	selectdeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err)
	}

	produto := Produto{}
	produtos := []Produto{}
	for selectdeTodosOsProdutos.Next() {
		err := selectdeTodosOsProdutos.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			panic(err)
		}
		produtos = append(produtos, produto)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
