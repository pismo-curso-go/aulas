// package main

// import (
// 	"fmt"
// 	"time"
// 	"weak"
// )

// type Data struct {
// 	Value string
// }

// func main() {
// 	data := &Data{Value: "Matheus"}
// 	// strongRef := data
// 	weakRef := weak.Make(data)

// 	data = nil

// 	// runtime.GC()
// 	time.Sleep(1 * time.Second)

// 	if weakRef.Value() == nil {
// 		fmt.Println("Dado sumiu")
// 	} else {
// 		fmt.Println("Dado ainda está lá")
// 	}

// 	// fmt.Println(strongRef.Value)
// }
