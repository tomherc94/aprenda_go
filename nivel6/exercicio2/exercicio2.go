/* - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
- Solução: https://play.golang.org/p/Tgv3wwxKV- */

package main

import "fmt"

func soma(valores ...int) int {
	total := 0
	for valor := range valores {
		total += valor
	}
	return total
}

func soma2(valores []int) int {
	total := 0
	for valor := range valores {
		total += valor
	}
	return total
}

func main() {

	valores := []int{3, 5, 6, 7, 8, 5}

	soma := soma(valores...)
	soma2 := soma2(valores)

	fmt.Println(soma, soma2)

}
