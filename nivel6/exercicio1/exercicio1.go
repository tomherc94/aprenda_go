/* - Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
- Solução: https://play.golang.org/p/rxJM5fgI-9
*/

package main

import "fmt"

func soma(x, y int) int {
	return x + y
}

func soma2(x, y int) (int, string) {
	return x + y, "Somou bonito!"
}

func main() {
	fmt.Println(soma(2, 4))
	fmt.Println(soma2(4, 9))
}
