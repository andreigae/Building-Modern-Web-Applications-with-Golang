package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	// se pide a la funcion changeUsingPointer que cambie el valor de myString
	// pasandole la direccion de memoria de myString (un puntero)
	changeUsingPointer(&myString)

	// Despues de llamar a la funcion, myString deberia ser Red
	log.Println("after func call myString is set to", myString)
}

// Esta funcion recibe un puntero a un string y lo cambia de Green a Red
func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
