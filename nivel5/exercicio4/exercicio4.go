/*- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
- Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG */

package main

import "fmt"

func main() {

	anonimo := struct {
		RA     int
		cursos []string
		notas  map[string]float32
	}{
		RA:     1231232,
		cursos: []string{"Administração", "TI"},
		notas: map[string]float32{
			"Administração": 10.0,
			"TI":            9.0,
		},
	}

	fmt.Println(anonimo)

}
