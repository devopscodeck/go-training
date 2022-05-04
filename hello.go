package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		// Best pratic with switch for this case
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibindoLogs()
		case 0:
			os.Exit(0) //informa que saímos do código com sucesso
		default:
			fmt.Println("O comando digitado não corresponde as opções disponíveis.")
			os.Exit(-1) //informa que algo saiu errado
		}
	}
}

func exibeIntroducao() {
	nome := "Administrator"
	var versao = 1.2

	fmt.Println("Hello", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")
}

func exibindoLogs() {
	fmt.Println("Exibindo logs..")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://alura.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com error: ", resp.StatusCode)

	}
}
