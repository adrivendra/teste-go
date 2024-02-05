package main

import (
	"fmt"
	"objetoOr/contas"
)

func main() {
	contaDoAdriano := contas.ContaCorrente{Titular: "Adriano", Saldo: 300}
	contaDoEdson := contas.ContaCorrente{Titular: "Edson", Saldo: 100}

	status := contaDoAdriano.Transferir(200, &contaDoEdson)

	fmt.Println(status)
	fmt.Println(contaDoAdriano)
	fmt.Println(contaDoEdson)

}
