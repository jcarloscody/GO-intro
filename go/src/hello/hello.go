package main

import "fmt"
import "reflect"
import "os"
import "net/http"

func main() {
	home()
	comando := leCom()
	controler(comando)
	requisicao()

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
	site := "https://www.youtube.com/watch?v=ipHf-wprgoQ&ab_channel=ADPBTemploCentral"
	res, _ := http.Get(site)
	fmt.Println("res::  ", res.Header)
	// fmt.Println("error  ::", error)
}

func doisReturn() (string, int) {
	return "ola", 0
}
