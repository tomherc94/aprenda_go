/*- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: https://play.golang.org/p/Gh81-d5tMi*/

package main

import "fmt"

func main() {

	slice := [][]string{
		{"Tomás", "Herculano", "Videogame"},
		{"Thaina", "Silva", "Assistir séries"},
		{"Rodrigo", "Germano", "Competição de tiro"},
	}

	for i, row := range slice {
		for _, value := range row {
			fmt.Println(i, value)
		}
	}

}
