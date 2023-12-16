/*- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
- Solução: https://play.golang.org/p/sA7NHpkCCg */

package main

import "fmt"

func somar() func() int {
	a := 10
	b := 20
	return func() int {
		return a + b
	}
}

func main() {

	teste1 := somar()
	fmt.Println(teste1())
}
