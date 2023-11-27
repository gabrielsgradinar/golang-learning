package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	// salva o valor do json em uma variável
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	// faz a serialização, pega o valor e entrega em algum lugar (stdout, arquivo)
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n":2, "s":200}`)
	var contaX Conta
	json.Unmarshal(jsonPuro, &contaX)

	println(contaX.Numero)
	println(contaX.Saldo)

}
