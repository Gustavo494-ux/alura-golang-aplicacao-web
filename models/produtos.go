package models

import (
	"alura-golang-aplicacao-web/db"
)

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

	selectdeTodosOsProdutos, err := db.Query("select * from produtos order by id")
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

func UPDATE(produto Produto) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id  = $5")
	if err != nil {
		panic(err.Error())
	}

	_, err = insereDadosNoBanco.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)
	if err != nil {
		panic(err.Error())
	}
}

func EditarProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	produtoDoBanco, err := db.Query("SELECT * FROM PRODUTOS WHERE id  = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}
	for produtoDoBanco.Next() {
		produtoDoBanco.Scan(
			&produtoParaAtualizar.Id,
			&produtoParaAtualizar.Nome,
			&produtoParaAtualizar.Descricao,
			&produtoParaAtualizar.Preco,
			&produtoParaAtualizar.Quantidade)

	}
	return produtoParaAtualizar
}

func DeletarProduto(idProduto int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deleteDadosNoBanco, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deleteDadosNoBanco.Exec(idProduto)
	if err != nil {
		panic(err.Error())
	}
}
