package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
	}

	delete(funcsPorLetra, "P")
	delete(funcsPorLetra["P"], "Pedro Junior")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Printf("Nome: %s, salário: %2.f\n", nome, salario)
		}
	}
}
