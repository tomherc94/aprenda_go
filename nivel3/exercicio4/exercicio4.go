/*- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/dIbfdiuw0ms */

package main

import "fmt"

func main() {

	ano := 1994
	atual := 2023

	for {
		fmt.Println(ano)
		ano++
		if ano > atual {
			break
		}
	}
}
