package main

import "fmt"
import "reflect"
import "os"
import "net/http"
import "time"

import "io"
import "io/ioutil"
import "bufio"
import "strings"
import "strconv"

const monitoramento = 5

func main() {
	lerFile()
	// for {
	// 	home()
	// 	comando := leCom()
	// 	controler(comando)
	// 	requisicao()
	// }

}

func home() {
	var name string = "jo"
	var nameWithou string
	sobreNome := "jo"
	var age float32 = 3.3
	fmt.Println("Hello word!!", name, nameWithou)
	fmt.Println("Hello word!!", sobreNome)
	fmt.Println("Hello word!!", age)
	fmt.Println("type ", reflect.TypeOf(age))

	fmt.Printf("1 iniciar, 2 logs, 0 sair")
}

func leCom() int {
	var comando int
	fmt.Scan(&comando) //& endereço da variavel comando
	return comando
}

func controler(comando int) {
	if comando == 1 {

	} else if comando == 2 {

	} else {

	}

	switch comando {
	case 1:
		fmt.Printf("1")
	case 2:
		fmt.Printf("2")
	default:
		os.Exit(0)
	}
}

func requisicao() {
	var arraySite [3]string
	arraySite[0] = "https://cursos.alura.com.br/formacao-go"

	sliceSites := []string{"https://www.youtube.com/watch?v=ipHf-wprgoQ&ab_channel=ADPBTemploCentral", "https://www.youtube.com/watch?v=ipHf-wprgoQ&ab_channel=ADPBTemploCentral", "https://www.youtube.com/watch?v=ipHf-wprgoQ&ab_channel=ADPBTemploCentral"} //simula um array dinamico
	sliceSites = append(sliceSites, "qq")
	fmt.Println("quantidade de itens", len(sliceSites))
	fmt.Println("capacidade", cap(sliceSites))

	site := "https://www.youtube.com/watch?v=ipHf-wprgoQ&ab_channel=ADPBTemploCentral"
	res, _ := http.Get(site)
	fmt.Println("res::  ", res.StatusCode)
	// fmt.Println("error  ::", error)

	for i := 0; i < 5; i++ {
		for i, v := range sliceSites {
			fmt.Println(i, v)
		}
		time.Sleep(monitoramento * time.Second)
	}

}

func doisReturn() (string, int) {
	return "ola", 0
}

func lerFile() []string {
	arquivo, _ := os.Open("sites.txt")
	reader := bufio.NewReader(arquivo)
	for {
		l, e := reader.ReadString('\n')
		l = strings.TrimSpace(l)
		fmt.Println(l)
		if e == io.EOF {
			break
		}
	}
	arquivo.Close()

	arquivoioutil, _ := ioutil.ReadFile("sites.txt")
	fmt.Println(string(arquivoioutil))

	log, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_CREATE|os.O_APPEND, 0666)
	log.WriteString("ola" + strconv.FormatBool(true) + time.Now().Local().String() + time.Now().Format("02/01/2006 15:04:05"))
	log.Close()

	arquivoioutil2, _ := ioutil.ReadFile("log.txt")
	fmt.Println(string(arquivoioutil2))

	return []string{}
}
