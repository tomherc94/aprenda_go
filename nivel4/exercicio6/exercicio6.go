/*- Crie uma slice usando make que possa conter todos os estados do Brasil.
    - Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
	"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
	"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice sem utilizar range.
- Solução: https://play.golang.org/p/cGYBphlyCE*/

package main

import "fmt"

func main() {

	estados := make([]string, 0, 10)

	estados = append(estados, "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins")

	fmt.Printf("Len: %v\nCap:%v\n\n", len(estados), cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}

}
