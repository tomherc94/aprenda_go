/*- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
- Solução: https://play.golang.org/p/zRBEooRvo4
*/

package main

import "fmt"

const (
	_ = 2023 + iota
	a
	b
	c
	d
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
