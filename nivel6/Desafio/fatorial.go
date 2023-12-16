package main

import "fmt"

func fatorial(valor int) int {
	if valor == 0 || valor == 1 {
		return 1
	}
	return valor * fatorial(valor-1)
}

func main() {

	fmt.Println(fatorial(5))

}
