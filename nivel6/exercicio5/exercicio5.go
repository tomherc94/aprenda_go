/* - Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
- Solução: https://play.golang.org/p/qLY-q3vffQ*/

package main

import (
	"fmt"
	"math"
)

type Figura interface {
	area() float64
}

type Quadrado struct {
	lado float64
}

func (q Quadrado) area() float64 {
	return math.Pow(q.lado, 2)
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

func info(f Figura) float64 {
	return f.area()
}

func main() {

	quadrado1 := Quadrado{2.0}
	circulo1 := Circulo{3.0}

	fmt.Printf("Área do quadrado1: %.2f\n", info(quadrado1))
	fmt.Printf("Área do circulo1: %.2f\n", info(circulo1))

}
