/* - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
- Solução: https://play.golang.org/p/b5Ua2kNN8a*/

package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	defer fmt.Println(7)
	fmt.Println(8)

}
