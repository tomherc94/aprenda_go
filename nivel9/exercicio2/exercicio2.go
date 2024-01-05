/* - Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
   - Crie um tipo para um struct chamado "pessoa"
   - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
   - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
   - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
   - Demonstre no seu código:
       - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
       - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"*/

package main

import "fmt"

type Humanos interface {
	Falar()
}

type Pessoa struct {
	Nome string
}

func (p *Pessoa) Falar() {
	fmt.Println(p.Nome + " está falando!")
}

func dizerAlgumaCoisa(h Humanos) {
	h.Falar()
}

func main() {
	pessoa1 := Pessoa{"Tomás"}

	dizerAlgumaCoisa(&pessoa1)
	//dizerAlgumaCoisa(pessoa1)

}
