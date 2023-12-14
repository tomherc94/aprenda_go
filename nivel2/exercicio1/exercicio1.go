/*- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
- Solução: https://play.golang.org/p/X7qm3aWSa6 */

package main

import "fmt"

func main() {
	n := 123
	fmt.Printf("%d\t%b\t%#x\n", n, n, n)
}
