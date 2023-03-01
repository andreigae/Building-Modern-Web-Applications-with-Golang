package main

import "log"

func main() {
	// the if statement
	var isTrue bool

	// declarar una varialbe y asignarle el valor de false
	isTrue = false

	// Si is true es igual a true
	if isTrue == true {
		// Imprimir en consola el valor de isTrue
		log.Println("isTrue is", isTrue)
	} else {
		// Si no es igual a true, imprimir en consola el valor de isTrue
		log.Println("isTrue is", isTrue)
	}

	// Esta forma de declarar una variable y asignarle un valor es mas comun
	// (Se declara y se agigna el valor en la misma linea)

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isTrue = false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		// do something
	} else if myNum == 101 || isTrue {
		// do something
	} else if myNum > 1000 && isTrue {
		// do something
	}

	// the switch statement

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "dog":
		log.Println("myVar is set to dog")

	case "horse":
		log.Println("myVar is set to horse")

	case "fish":
		log.Println("myVar is set to fish")

	default:
		log.Println("myVar is something else")
	}
}
