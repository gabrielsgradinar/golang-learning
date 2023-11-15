package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Println("Desativou a empresa")
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O client %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 25,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	// gabriel.Desativar()
	Desativacao(gabriel)
	Desativacao(minhaEmpresa)

	gabriel.Endereco.Cidade = "São Paulo"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", gabriel.Nome, gabriel.Idade, gabriel.Ativo)

}
