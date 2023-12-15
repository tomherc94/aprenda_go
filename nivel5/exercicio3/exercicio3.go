/* - Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
- Solução: https://play.golang.org/p/3eoGb0kxzT*/

package main

import "fmt"

type Veiculo struct {
	portas int
	cor    string
}

type Caminhonete struct {
	Veiculo
	tracao4x4 bool
}

type Sedan struct {
	Veiculo
	modeloLuxo bool
}

func main() {

	caminhonete1 := Caminhonete{
		Veiculo: Veiculo{
			portas: 3,
			cor:    "azul",
		},
		tracao4x4: true,
	}

	sedan1 := Sedan{
		Veiculo: Veiculo{
			portas: 5,
			cor:    "vermelho",
		},
		modeloLuxo: false,
	}

	fmt.Println(caminhonete1.portas, caminhonete1.cor, caminhonete1.tracao4x4)
	fmt.Println(sedan1.portas, sedan1.cor, sedan1.modeloLuxo)

	fmt.Println(caminhonete1.cor, sedan1.portas)
}
