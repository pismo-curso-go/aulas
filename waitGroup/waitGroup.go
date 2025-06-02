// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func escrever(texto string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(texto)
// 		time.Sleep(time.Second)
// 	}
// }

// func escrever2(texto string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(texto)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {

// 	var waitGroup sync.WaitGroup

// 	waitGroup.Add(3) // Crédito 3

// 	fmt.Println("Disparando goroutine 01")
// 	go func() {
// 		escrever("Escrever 1")
// 		waitGroup.Done() // -1
// 	}()

// 	fmt.Println("Disparando goroutine 02")
// 	go func() {
// 		escrever("Escrever 2")
// 		waitGroup.Done() // -1
// 	}()

// 	fmt.Println("Disparando goroutine 03")
// 	go func() {
// 		escrever("Escrever 3")
// 		waitGroup.Done() // -1
// 	}()

// 	waitGroup.Wait() // pare aqui - 0

// 	fmt.Println("Continua execução")
// }
