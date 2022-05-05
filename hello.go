package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3 //quantidade de sites monitorados
const delay = 5          // tempo de espera para requisitar os sites novamente

func main() {
	exibeIntroducao()
	fmt.Println("")
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
	//a linha 59 refere-se a um slice (é um array por tras do capô)
	//	sites := []string{"https://alura.com.br", "https://alura.com.br/dashboards", "https://alura.com.br/ddddd", "https://cursos.alura.com.br/course/golang/task/27970"}
	sites := leSitesDoArquivo()

	//testa de 5 em 5 segundos
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("testando site:", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Print("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com error: ", resp.StatusCode)
	}

}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
