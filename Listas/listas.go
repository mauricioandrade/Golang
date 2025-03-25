
package main

import "fmt"

func main() {
	
	// Array tamanho fixo
	var array[2]string
	array[0] = "Mauricio"
	array[1] = "Andrade"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2,3,5,7,11,13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3])


	//Slices -- !! LEMBRANDO QUE SEMPRE PRECISAMOS USAR O MAKE! 
	slice := make([]string,5)
	slice[0] ="Mauricio"
	slice[1] = "Andrade"
	fmt.Println(slice[0],slice[1])
	fmt.Println(slice)
	fmt.Println(len(slice))

	numPares := []int{2,4,6,8,10,12}
	fmt.Println(numPares)
	

	numPares = append(numPares, 14) // adcionar valor no array
	fmt.Println(numPares)
}

//LISTAS

//1- arrays ou slice : homogeneos
// todos os elementos vão ser do mesmo tipo
// [1,2,3,4,5,6] - []int
//["Mauricio","Andrade"] - []String

//2 - maps: heterogeneos
//pode misturar tipos
// estrutura chave - valor
//[key] = value
// chave tem um tipo, e o valor pode ter outro 
//map[string]int
//{"Mauricio": 32, "Tiquinha": 8}
//map[string]string
//{"Mauricio": "Andrade"}

//ARRAY
// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com indice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// Por conta do tamanho fixo, não é tão usado. Só em casos especificos

//Slice
// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com indice: a[0],a[1]
// Função embutida len() retorna o tamanho do slice
// Função append() usada para adcionar valores no slice
