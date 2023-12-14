/*- Crie um programa que utilize a declaração switch, sem switch statement especificado.
- Solução: https://play.golang.org/p/TyIGp4Hi8B */

package main

import "fmt"

func main() {
	diaSemana := 74

	switch {
	case diaSemana == 1:
		fmt.Println("Domingo")
	case diaSemana == 2:
		fmt.Println("Segunda-feira")
	case diaSemana == 3:
		fmt.Println("Terça-feira")
	case diaSemana == 4:
		fmt.Println("Quarta-feira")
	case diaSemana == 5:
		fmt.Println("Quinta-feira")
	case diaSemana == 6:
		fmt.Println("Sexta-feira")
	case diaSemana == 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Dia inválido!")
	}
}
