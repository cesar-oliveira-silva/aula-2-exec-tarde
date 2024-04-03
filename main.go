package main

import (
	"fmt"
)

// exercicio 1
type aluno struct {
	nome      string
	sobrenome string
	rg        int
	dataAdm   string
}

var listaAluno []aluno

//exercicio 2

type loja struct {
	nome     string
	produtos []produto
}

type produto struct {
	nome  string
	tipo  string
	preco float64
}

type iproduto interface {
	calcularCusto()
}

type iloja interface {
	total()
	adicionar()
}

func main() {

	//exercicio 1
	var novoAluno1 aluno
	novoAluno1.nome = "Cesar"
	novoAluno1.sobrenome = "Silva"
	novoAluno1.rg = 123456
	novoAluno1.dataAdm = "01-01-2024"
	cadastraAluno(novoAluno1)

	listaAlunos()

	// exercicio 2
	pequeno1 := novoProduto("celular", "pequeno", 2000.0)
	pequeno2 := novoProduto("mouse", "pequeno", 100.0)
	medio1 := novoProduto("monitor", "medio", 1000.0)
	medio2 := novoProduto("gabinete", "medio", 200.0)
	grande1 := novoProduto("mesa", "grande", 200.0)
	grande2 := novoProduto("cadeira", "grande", 200.0)

	lojao := novaLoja("Lojao", []produto{})

	lojao.adicionar(pequeno1)
	lojao.adicionar(pequeno2)
	lojao.adicionar(medio1)
	lojao.adicionar(medio2)
	lojao.adicionar(grande1)
	lojao.adicionar(grande2)

	fmt.Printf("Total da Loja: %v \n", lojao.total())

}

// exercicio 1
func cadastraAluno(alun aluno) {
	listaAluno = append(listaAluno, alun)
}

func listaAlunos() {
	for _, alun := range listaAluno {
		fmt.Printf("Nome Aluno: %v, Sobrenome: %v, Rg: %v, Data Admissao: %v \n", alun.nome, alun.sobrenome, alun.rg, alun.dataAdm)
	}
}

//exercicio 2

func (p produto) calcularCusto() float64 {

	switch {
	case p.tipo == "pequeno":
		return p.preco

	case p.tipo == "medio":
		return p.preco + (p.preco * 0.03)

	case p.tipo == "grande":
		return p.preco + (p.preco * 0.06) + 2500

	}
	return 0
}

func novoProduto(name string, tipe string, value float64) produto {
	fmt.Printf("Criando o Produto %v \n", name)
	return produto{nome: name, tipo: tipe, preco: value}
}

func novaLoja(name string, items []produto) loja {
	fmt.Printf("Criando a Loja %v \n", name)
	return loja{nome: name, produtos: items}
}

func (l loja) total() float64 {
	acc := 0.0
	for _, item := range l.produtos {
		acc += item.calcularCusto()
	}
	return acc
}

func (l *loja) adicionar(p produto) {
	fmt.Printf("Adicionando o produto %v a loja %v \n", p.nome, l.nome)
	l.produtos = append(l.produtos, p)

}
