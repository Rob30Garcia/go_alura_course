package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	ler_sites_do_arquivo()
	exibe_introducao()
	registra_log("site-falso", false)
	for {
		exibe_menu()
		comando := ler_comando()

		switch comando {
		case 1:
			iniciando_monitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibe_introducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibe_menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func ler_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comand é", &comando)
	fmt.Println("O comando escolhido foi", comando)
	fmt.Println()
	return comando
}

func iniciando_monitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://github.com/", "https://www.alura.com.br/", "https://jovemnerd.com.br/"}
	sites := ler_sites_do_arquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site:", site)
			test_site(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println()
}

func test_site(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registra_log(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:", res.StatusCode)
		registra_log(site, false)
	}
}

func ler_sites_do_arquivo() []string {
	var sites []string

	// file, err := os.ReadFile("sites.txt")
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(file)

	for {
		line, err := leitor.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registra_log(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
