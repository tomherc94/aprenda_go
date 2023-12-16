/* - Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
- Solução: https://play.golang.org/p/GBZcnu0Ajp*/

package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p *Pessoa) toString() string {
	saida := fmt.Sprintf("Nome completo: %v %v\nIdade: %v", p.Nome, p.Sobrenome, p.Idade)
	return saida
}
func main() {

	pessoa := Pessoa{"Tomas", "Herculano", 29}

	fmt.Println(pessoa.toString())

}
