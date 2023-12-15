/*- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: https://play.golang.org/p/nD3TW8VQmH*/

package main

import "fmt"

func main() {

	meumap := map[string]string{
		"Herculano_Tomas": "jogar videogame",
		"Silva_Thaina":    "Assistir séries",
		"Germano_Rodrigo": "Competição de tiro",
	}

	for key, value := range meumap {
		fmt.Println(key, value)
	}

}
