package models

import "fmt"
import "loja/db"

type Produto struct {
	Id int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	fmt.Println("Conectando no banco para buscar dados.")
	db := db.ConectaBancoDados()
	fmt.Println("Conectou.")


	selectTodosProdutos, err := db.Query("select * from produtos order by idprodutos asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	fmt.Println("Fechou a conexao.")

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int){
	fmt.Println("Conectando no banco para gravar dados.")
	db := db.ConectaBancoDados()
	fmt.Println("Conectou.")
	
	insereDadosBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values (?, ?, ?, ?)")
	if err != nil{
		panic(err.Error())
	}

	insereDadosBanco.Exec(nome, descricao, preco, quantidade)
	fmt.Println("Inseriu.")

	defer db.Close()
	fmt.Println("Fechou a conexao.")
}

func DeletaProduto(id string){
	fmt.Println("Conectando no banco para apagar dados.")
	db := db.ConectaBancoDados()
	fmt.Println("Conectou.")

	deletarProduto, err := db.Prepare("delete from produtos where idprodutos=?")
	if err != nil{
		panic(err.Error())
	}

	deletarProduto.Exec(id)
	fmt.Println("Apagou.")

	defer db.Close()
	fmt.Println("Fechou a conexao.")
}

func EditaPRoduto(id string) Produto {
	fmt.Println("Conectando no banco para buscar um produto.")
	db := db.ConectaBancoDados()
	fmt.Println("Conectou.")

	buscarProduto, err := db.Query("select * from produtos where idprodutos=?", id)
	if err != nil{
		panic(err.Error())
	}
	
	produtoParaAtualizar := Produto{}
	for buscarProduto.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = buscarProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil{
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Preco = preco
	}
	
	fmt.Println("Pegou o produto")

	defer db.Close()
	fmt.Println("Fechou a conexao.")

	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int){
	fmt.Println("Conectando no banco para atualizar um produto.")
	db := db.ConectaBancoDados()
	fmt.Println("Conectou.")

	atualizaProduto, err := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where idprodutos=?")
	if err != nil{
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	fmt.Println("Atualizou.")

	defer db.Close()
	fmt.Println("Fechou a conexao.")
}