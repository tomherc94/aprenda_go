/*- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
- Demonstre estes valores.
- Solução: https://play.golang.org/p/BMYEch6_s8 */

package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10.0
	b := 4.0
	soma := a + b
	subtracao := a - b
	divisao := a / b
	multiplicacao := a * b
	potenciacao := math.Pow(a, b)

	fmt.Printf("Soma: %v\n", soma)
	fmt.Printf("Subtração: %v\n", subtracao)
	fmt.Printf("Divisão: %v\n", divisao)
	fmt.Printf("Multiplicação: %v\n", multiplicacao)
	fmt.Printf("Potenciação: %v\n", potenciacao)
}
