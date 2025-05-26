// package main

// import "fmt"

// func main() {
// 	slice := make([]int, 10)

// 	slice[5] = 20
// 	fmt.Println(slice)

// 	slice = make([]int, 10, 20)
// 	fmt.Println(slice, len(slice), cap(slice))

// 	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
// 	fmt.Println(slice, len(slice), cap(slice))

// 	slice = append(slice, 1)
// 	fmt.Println(slice, len(slice), cap(slice))

// 	slice4 := make([]int, 10, 20)

// 	slice2 := make([]int, 10, 20)
// 	slice3 := append(slice2, 10, 20, 30)

// 	fmt.Println(slice2, slice3)

// 	slice2[0] = 3
// 	slice3[0] = 8

// 	fmt.Println(slice2, slice3, slice4)

// }
