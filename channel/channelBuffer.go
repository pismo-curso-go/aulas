package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1 // [1] [] []
	ch <- 2 // [1] [2] []
	ch <- 3 // [1] [2] [3]
	ch <- 4 // [2] [3] [4]
	ch <- 5 // [3] [4] [5]
	ch <- 6
	fmt.Println("Executou!")
}

func main() {
	canal := make(chan int, 3) // [][][]
	go rotina(canal)           // Async e independente

	time.Sleep(time.Second * 5)
	fmt.Println(<-canal) // 1
}
