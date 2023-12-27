package controllers

import (
	"alura-golang-aplicacao-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.BuscarTodosProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var produto models.Produto

		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		var err error
		produto.Preco, err = strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		produto.Quantidade, err = strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		models.CriarNovoProduto(produto)

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	models.DeletarProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
