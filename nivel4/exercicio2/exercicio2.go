/*- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
- Solução: https://play.golang.org/p/ST3TKusuOd */

package main

import "fmt"

func main() {

	slice := make([]int, 0)

	for i := 0; i < 10; i++ {
		slice = append(slice, i+1000)
	}

	for i, value := range slice {
		fmt.Println(i, value)
	}

	fmt.Printf("Tipo: %T\n", slice)

}
