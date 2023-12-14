/*- Utilizando a solução anterior, adicione as opções else if e else.
- Solução: https://play.golang.org/p/_cO7E-yL0o */

package main

import "fmt"

func main() {
	nota := 1.1

	if nota >= 6.0 {
		fmt.Println("Aprovado!")
	} else if nota >= 3.0 {
		fmt.Println("Recuperação!")
	} else {
		fmt.Println("Reprovado!")
	}
}
