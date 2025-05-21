package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

// "Método": função com receiver (receptor)
func (produto Produto) precoComDesconto() float64 {
	return produto.preco * (1 - produto.desconto)
}

func main() {
	var produto1 Produto

	produto1 = Produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	produto2 := Produto{
		nome:     "Caneta",
		preco:    2.00,
		desconto: 0.5,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produto2, produto2.precoComDesconto())
}
