/* - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: https://play.golang.org/p/GLK11Q1_x8y*/

package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresDeSorvete []string
}

func main() {

	meumap := make(map[string]Pessoa, 0)

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

	meumap["Herculano"] = pessoa1
	meumap["Silva"] = pessoa2

	for _, p := range meumap {
		fmt.Println(p.nome, p.sobrenome)
		for _, sabor := range p.saboresDeSorvete {
			fmt.Println("\t", sabor)
		}
		fmt.Println()
	}

}
