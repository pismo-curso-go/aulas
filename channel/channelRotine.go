// package main

// import (
// 	"fmt"
// 	"time"
// )

// func doisTresQuatroVezes(base int, ch chan int) {
// 	time.Sleep(time.Second)
// 	ch <- 2 * base
// 	time.Sleep(time.Second)
// 	ch <- 3 * base
// 	time.Sleep(3 * time.Second)
// 	ch <- 4 * base
// }

// func main() {
// 	canal := make(chan int)
// 	go doisTresQuatroVezes(2, canal)
// 	fmt.Println("A")

// 	a, b := <-canal, <-canal // sincronismo
// 	fmt.Println("B")
// 	fmt.Println(a, b)

// 	fmt.Println(<-canal)
// }
