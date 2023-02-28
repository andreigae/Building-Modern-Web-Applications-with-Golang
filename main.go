package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	// Guardar LOS RESULTADOS de la función en variables (whatWasSaid, theOtherThingThatWasSaid)
	// GO permite retornar multiples valores
	// En este caso, la función retorna dos strings
	// Lo que hace esta linea realmente es, ejecutar la funcion saySometing()
	// y sea lo que sea que devuelve esa función, lo guarda en las variables

	// PD: La funcion esta tipificada para retornar dos strings por lo que
	// no se puede guardar en variables otro tipo de informacion, ya que de por si
	// la funcion esta obligada a devolver 2 strings y ante cualqueir otro error
	// no copila directamente el programa

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
