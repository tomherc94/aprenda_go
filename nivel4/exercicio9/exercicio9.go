/*- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
- Solução: https://play.golang.org/p/3fcvHlt8Lm*/

package main

import "fmt"

func main() {

	meumap := map[string]string{
		"Herculano_Tomas": "jogar videogame",
		"Silva_Thaina":    "Assistir séries",
		"Germano_Rodrigo": "Competição de tiro",
	}

	meumap["Silva_Bruno"] = "Lutar Jiu-jitsu"

	for key, value := range meumap {
		fmt.Println(key, value)
	}

}
