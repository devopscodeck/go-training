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
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	// Condictional convenctional with if and else:
	// 	if comando == 1 {
	// 		fmt.Println("Monitorando...")
	// 	} else if comando == 2 {
	// 		fmt.Printf("Exibindo logs..")
	// 	} else if comando == 0 {
	// 		fmt.Printf("Saindo do programa")
	// 	} else {
	// 		fmt.Printf("Comando não existe como opção")
	// 	}
	// }

	// Best pratic with switch for this case
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs..")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("O comando digitado não corresponde as opções disponíveis.")
	}
}
