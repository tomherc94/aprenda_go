/* - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
- Solução: https://play.golang.org/p/Pyp7vmTJfY*/

package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresDeSorvete []string
}

var lista []Pessoa

func main() {

	pessoa1 := Pessoa{
		nome:             "Tomas",
		sobrenome:        "Herculano",
		saboresDeSorvete: []string{"Flocos", "Chocolate"},
	}

	pessoa2 := Pessoa{
		nome:             "Thaina",
		sobrenome:        "Silva",
		saboresDeSorvete: []string{"Napolitano", "Coco"},
	}

	lista = append(lista, pessoa1, pessoa2)

	for _, p := range lista {
		fmt.Println(p.nome, p.sobrenome)
		for _, sabor := range p.saboresDeSorvete {
			fmt.Println("\t", sabor)
		}
		fmt.Println()
	}

}
