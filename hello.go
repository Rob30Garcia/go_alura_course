package main

import (
	"fmt"
	"os"
)

func main() {
	exibe_introducao()

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	comando := ler_comando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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

func exibe_introducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func ler_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comand é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}
