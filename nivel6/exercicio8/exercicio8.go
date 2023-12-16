/*- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
- Solução: https://play.golang.org/p/A74rufv6Rs */

package main

import "fmt"

func geradorDeFuncoes(operador string) func(a, b int) int {
	switch {
	case operador == "+":
		return func(a, b int) int {
			return a + b
		}
	case operador == "-":
		return func(a, b int) int {
			return a - b
		}
	case operador == "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		return nil
	}
}

func main() {

	somador := geradorDeFuncoes("+")
	fmt.Println(somador(10, 20))

	subtrador := geradorDeFuncoes("-")
	fmt.Println(subtrador(10, 20))

	multiplicador := geradorDeFuncoes("*")
	fmt.Println(multiplicador(10, 20))

}
