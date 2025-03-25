package main

import (
	"fmt"

)

func main() {

	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)
	subtrairValores := subtrair(42, 13)
	fmt.Println(subtrairValores)
	idadeCalculo := idade(2025, 1992)
	fmt.Println(idadeCalculo)
	multiplicarValores := multiplicar(30, 30)
	fmt.Println(multiplicarValores)
	decimaisValores := decimais(12.3, 14.6)
	fmt.Println(decimaisValores)

	nome1, sobrenome := printaNome("Mauricio", "Andrade")
	fmt.Println(nome1)
	fmt.Println(sobrenome)

	
	fmt.Println(printaDiasUteis())
	fmt.Println(dividirValores())

	//scanner go
	var nome string
	var idade int
	var altura float64
	fmt.Scan(&nome)
	fmt.Scan(&idade)
	fmt.Scan(&altura)
	fmt.Printf(nome, idade, altura)
}

// função começando com letra minuscula:
// FUNÇÃO É PRIVADA 
// ELA SÓ PODE SER UTILIZADA NO PACOTE MAIN ex printaNome

// função começando com letra maiscula:
// FUNÇÃO É PUBLICA
// PODE SER CHAMADA EM OUTRO PACOTE ex PrintaNome
// como utilizar ela fora do pacote, ex main.PrintaNome(nome)



func printaNome(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

func printaDiasUteis() string {
	diasUteis := "Segunda, Terça, Quarta, Quinta, Sexta"
	return diasUteis
}

func dividirValores() int{
	dividir := 100/25 
	return dividir
}


func soma(x int, y int) int {
	return x + y
}

func multiplicar(x int, y int) int {
	return x * y
}

func decimais(x float64, y float64) float64 {
	return x + y
}

func subtrair(x int, y int) int {
	return x - y
}

func idade(anoatual int, anonascimento int) int {
	return anoatual - anonascimento
}
