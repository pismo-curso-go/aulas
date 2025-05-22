package main

import "fmt"

func main() {

	nomes := [...]string{"Marcos", "João", "Maria", "José", "Lucas"}

	for index := range nomes {
		fmt.Println(index)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for i, nome := range nomes {
		fmt.Printf("%d) %s\n", i, nome)
	}
}
