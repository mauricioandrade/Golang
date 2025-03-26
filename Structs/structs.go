package main

import "fmt"

	//structs
	// forma de criar sua propria estrutura de dados
	// personalizar de acordo com a sua necessidade
	// podemos usar vários tipos diferentes

//type <nome da estrutura> struct { <campos }
// não precisamos setar todos os campos na estrutura.

type Pessoa struct {
	Nome  string
	Sobrenome string
	idade int
	anoNasc int
}

type Profissao struct{
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Mauricio","Andrade",32,1992}) // forma de fazer a struct
	fmt.Println(Pessoa{Nome: "Mauricio", Sobrenome: "Andrade", idade: 32, anoNasc: 1992}) // melhor forma de fazer, passando o nome dos campos juntos

	p1 := Pessoa{Nome: "Mauricio", Sobrenome: "Andrade", idade: 32,anoNasc: 1992}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Sobrenome)
	fmt.Println(p1.idade)
	fmt.Println(p1.anoNasc)

	p1.idade = 33
	fmt.Println(p1.idade)

	 p2 := Pessoa{Nome: "Tiquinha", Sobrenome: "Cachorro", idade: 8, anoNasc: 2017 }

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1,p2)
	fmt.Println(pessoas)

	// // struct + maps
	// alunos := map[string][]Pessoa{}
	// alunos["programação"] = pessoas
	// fmt.Println(alunos)

	var alunos = map[string][]Pessoa{
		"Programação": {{Nome: "Mauricio", Sobrenome: "Andrade", idade: 32, anoNasc: 1992}},
		"Engenharia": {{Nome: "Tiquinha", Sobrenome: "Cachorro", idade: 8, anoNasc: 2017}},
	}
	fmt.Println(alunos)

// struct herdando campos de outra struct
prof := Profissao{p2, "dev"}
fmt.Println(prof)
fmt.Println(prof.Nome)
fmt.Println(prof.Sobrenome)
fmt.Println(prof.idade)
fmt.Println(prof.anoNasc)
fmt.Println(prof.Tipo)




}
