// package main

// import "fmt"

// type Sucavel interface {
// 	Beber() string
// 	Derramar() string
// }

// type SucoDeLaranja struct{}
// type SucoDeUva struct{}
// type SucoDeAbacaxi struct{}

// func (s SucoDeAbacaxi) Bebe() string {
// 	return "Tomando suco de Abacaxi"
// }

// func (s SucoDeLaranja) Beber() string {
// 	return "Tomando suco de laranja"
// }
// func (s SucoDeLaranja) Derramar() string {
// 	return "Derramando suco de laranja"
// }

// func (s SucoDeUva) Beber() string {
// 	return "Tomando suco de uva"
// }

// func ServirCopoAbacaxi(abacaxi SucoDeAbacaxi) {
// 	fmt.Println(abacaxi.Bebe())
// }

// func ServirCopo(bebida Sucavel) {
// 	fmt.Println(bebida.Beber())
// }

// func main() {

// 	laranja := SucoDeLaranja{}
// 	uva := SucoDeUva{}

// 	ServirCopo(laranja)
// 	ServirCopo(uva)
// }
