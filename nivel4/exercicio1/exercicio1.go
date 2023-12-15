/*- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
- Solução: https://play.golang.org/p/tpWDzzlO2l */

package main

import "fmt"

var array [5]int

func main() {

	for i := 0; i < len(array); i++ {
		array[i] = i + 1000
	}

	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Printf("Tipo: %T\n", array)

}
