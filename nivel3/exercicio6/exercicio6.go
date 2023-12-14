/*- Crie um programa que demonstre o funcionamento da declaração if.
- Solução: https://play.golang.org/p/OGPgTJZbiy */

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
