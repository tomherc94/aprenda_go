/*- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
- Solução: https://play.golang.org/p/1aPXVeR1mf */

package main

import "fmt"

func main() {

	slice := make([]int, 0)

	for i := 0; i < 10; i++ {
		slice = append(slice, i+1000)
	}
	fmt.Println(slice[0:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])

}
