package main

import "fmt"

// loops
// laços de repetições
// repetir tarefas

func main() {

	// i++ -> soma 1
	// i-- -> subtraindo 1

	// for range
	// frutas := []string{"Laranja", "Maça", "Banana", "Uva", "Kiwi"}
	// for i, fruta := range frutas {
	// 	fmt.Println("Fruta",fruta, "indice", i)
	// }

	champs := []string{"Kai'sa", "Caitlyn", "Vayne", "Lucian", "Ezreal"}
	for _, champ := range champs {
		fmt.Println("Os Champs são :", champ)
	}

	// sum := 0
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("Sum: ", sum)
	// 	fmt.Println("Indice: ", i)
	// 	sum += i

	// contando até 10
		for i := 0; i <=10; i++{
			fmt.Println(i)
		}

		// é a mesma coisa que: sum = sum + i
		// sum -= i -> sum = sum - i

		// é como se ao final do loop fosse adcionado 1 ao indice i
		// i++
		// i = i + 1

	}

