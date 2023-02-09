package main

import (
	"fmt"

	"github.com/rromulos/go-conta/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	conta1 := contas.ContaPoupanca{}
	conta1.Depositar(100)
	PagarBoleto(&conta1, 60)

	fmt.Println(conta1.ObterSaldo())

	conta2 := contas.ContaCorrente{}
	conta2.Depositar(500)
	PagarBoleto(&conta2, 1000)

	fmt.Println(conta2.ObterSaldo())

}