package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Administrator"
	var versao = 1.2
	idade := 25

	fmt.Println("Hello", nome, "sua idade é:", idade)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("O tipo da variável nome, é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável versao, é", reflect.TypeOf(versao))
}
