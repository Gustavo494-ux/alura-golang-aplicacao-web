package models

import "alura-golang-aplicacao-web/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
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
	return produtos
}

func CriarNovoProduto(produto Produto) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome,descricao, preco, quantidade) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insereDadosNoBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	if err != nil {
		panic(err.Error())
	}
}
