/*- Desafio surpresa!
- Format printing:
    - Decimal       %d
    - Hexadecimal   %#x
    - Unicode       %#U
    - Tab           \t
    - Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
- Solução: https://play.golang.org/p/REm2WHyzzz*/

package main

import "fmt"

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("\tDecimal: %d\tHexadecimal: %#x\tUnicode: %#U\n", i, i, i)
	}
}
