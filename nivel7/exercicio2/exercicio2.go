/*- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
- Solução: https://play.golang.org/p/qiYp9leJcn
*/

package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
}

func mudeMe(pessoa *Pessoa) {
	(*pessoa).Nome = "Nome alterado!"
	pessoa.Sobrenome = "Sobrenome alterado!"
}

func main() {
	pessoa1 := Pessoa{"Tomas", "Herculano"}
	mudeMe(&pessoa1)
	fmt.Println(pessoa1.Nome, pessoa1.Sobrenome)
}
