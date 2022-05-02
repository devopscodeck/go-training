package main

import (
	"fmt"
)

func main() {
	var nome string = "Administrator"
	var versao = 1.2

	fmt.Println("Hello", nome)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
