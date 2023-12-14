/*- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
- Solução: https://play.golang.org/p/RkpqPpRWuo */

package main

import "fmt"

func main() {

	a := `Esta é
	uma         RAW
	STRING
									!!!`

	fmt.Println(a)
}
