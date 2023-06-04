package main

import "fmt"
import "reflect"
import "os"

func main() {
	home()
	comando := leCom()
	controler(comando)

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
