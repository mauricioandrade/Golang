package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Mauricio"] = 32
	idade["Tiquinha"] = 8
	fmt.Println(idade)
	fmt.Println(idade["Mauricio"])
	fmt.Println(idade["Tiquinha"])

	anoNasc := map[string]int{
		"Mauricio": 1992,
		"Tiquinha": 2017,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Mauricio"])
	fmt.Println(anoNasc["Tiquinha"])
	anoNasc["GolangDoZero"] = 2025
	fmt.Println(anoNasc)


	alturaPessoa :=map[string]float64{
		"Mauricio": 1.96,
		"Vera": 1.70,

	}
	fmt.Println(alturaPessoa)
		fmt.Println(alturaPessoa["Mauricio"])
		fmt.Println(alturaPessoa["Vera"])
	}


