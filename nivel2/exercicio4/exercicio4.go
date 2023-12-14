/*- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
- Solução: https://play.golang.org/p/IiwgT0v3Mp */

package main

import "fmt"

func main() {

	a := 2
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)

}
