/*- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do
tipo string com identificador "esporteFavorito".
- Solução: https://play.golang.org/p/4-iTPZwfEz */

package main

import "fmt"

func main() {
	esporteFavorito := "basquete"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Neymar é top!")
	case "natação":
		fmt.Println("Gustavo Borges é top!")
	case "basquete":
		fmt.Println("Varejão é top!")
	default:
		fmt.Println("Esporte inválido!")
	}
}
