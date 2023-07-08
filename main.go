package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func showGreeting() {
	nome := "Breno"
	versao := 1.1

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)

}

func showOptionsMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func readOption() int {
	var option int
	fmt.Print("Selecione uma opção: ")
	fmt.Scan(&option)

	return option
}

func verifySiteStatus(url string) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		fmt.Printf("Site: %s foi carregado com sucesso\n", url)
		writeLogsFile(url, true)
	} else {
		fmt.Printf("Não foi possivel estabelecer uma conexão. Status Code: %d\n", res.StatusCode)
		writeLogsFile(url, false)
	}
}

func readUrlFile() []string {
	var urlArr []string
	file, err := os.Open("url.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		urlArr = append(urlArr, line)
	}

	return urlArr
}

func writeLogsFile(url string, online bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	const timeFormat = "02/01/2006 15:04:05"

	file.WriteString(time.Now().Format(timeFormat) + " - " + url + " - online: " + strconv.FormatBool(online) + "\n")
}

func printLogs() {
	fmt.Println("Exibindo logs...")
	file, err := os.ReadFile("logs.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))
}

const monitoring int = 3
const delay = 5

func startMonitoring() {
	fmt.Println("Monitorando...")

	urlArr := readUrlFile()

	for i := 0; i < monitoring; i++ {
		for _, url := range urlArr {
			verifySiteStatus(url)
		}

		time.Sleep(delay * time.Second)
	}
}

func main() {
	showGreeting()

	for {
		showOptionsMenu()

		option := readOption()

		fmt.Println("O valor da variável comando é:", option)

		switch option {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}

	}
}
