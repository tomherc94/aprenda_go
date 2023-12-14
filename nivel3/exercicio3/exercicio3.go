/*- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/qnFjiDJzLor*/

package main

import "fmt"

func main() {

	ano := 1994
	atual := 2023

	for ano <= atual {
		fmt.Println(ano)
		ano++
	}
}
