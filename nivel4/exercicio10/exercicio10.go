/*- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.*/

package main

import "fmt"

func main() {

	meumap := map[string]string{
		"Herculano_Tomas": "jogar videogame",
		"Silva_Thaina":    "Assistir séries",
		"Germano_Rodrigo": "Competição de tiro",
	}

	meumap["Silva_Bruno"] = "Lutar Jiu-jitsu"

	delete(meumap, "Germano_Rodrigo")

	for key, value := range meumap {
		fmt.Println(key, value)
	}

}
