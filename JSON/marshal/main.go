// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type cachorro struct {
// 	NomeCanino string `json:"nome_canino"`
// 	Raca       string `json:"raca"`
// 	Idade      uint   `json:"idade"`
// }

// func main() {
// 	c := cachorro{"Rex", "Dálmata", 3}

// 	cachorroEmJSON, erro := json.Marshal(c)
// 	if erro != nil {
// 		log.Fatal(erro)
// 	}
// 	fmt.Println(cachorroEmJSON)
// 	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

// 	c2 := map[string]string{
// 		"nome": "Toby",
// 		"raca": "Poodle",
// 	}

// 	cachorro2EmJSON, erro := json.Marshal(c2)
// 	if erro != nil {
// 		log.Fatal(erro)
// 	}
// 	fmt.Println(cachorro2EmJSON)
// 	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
// }
