/* - Callback: passe uma função como argumento a outra função.
- Solução: https://play.golang.org/p/2epLD_Yyap*/

package main

import "fmt"

func somar(x func(x, y int) int) {
	fmt.Println(x(10, 20))
}

func main() {

	somar(func(a, b int) int {
		return a + b
	})

}
