package main

import "fmt"
import "reflect"

func main() {
	var name string = "jo"
	var nameWithou string
	sobreNome := "jo"
	var age float32 = 3.3
	fmt.Println("Hello word!!", name, nameWithou)
	fmt.Println("Hello word!!", sobreNome)
	fmt.Println("Hello word!!", age)
	fmt.Println("type ", reflect.TypeOf(age))

	fmt.Printf("1 iniciar, 2 logs, 0 sair")
	var comando int
	fmt.Scan(&comando) //& endereço da variavel comando

	if comando == 1 {

	} else if comando == 2 {

	} else {

	}

}
