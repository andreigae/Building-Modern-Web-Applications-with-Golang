package main

import "log"

type myStruct struct {
	FirstName string
}

// Se crea una funcion printFirstName() dentro de la estructura myStruct
// esta funcion regresa el valor del campo FirstName de la estructura
func (myStr *myStruct) printFirstName() string {
	return myStr.FirstName
}

/*
PD: Normalmente una funcion se declara como:

func printFirstName() string {

}

Pero eso es una funcion sin mas, que no pertenece a ninguna estructura.

Podriamos hacer que llamando a esta funcion se imprima el valor del campo de interes
pero no es lo mismo que hacerlo dentro de la estructura, ya que la funcion necesita recibir
la estructura como parametro, y eso no es lo que queremos

Lo que quieremos es que la funcion sea parte de la estructura, y que no necesite recibir nada
=> Que sea parte de la estructura y que devuelva su campo FirstName.

Por eso  (myStr *myStruct), es lo que se denomina un reciver,
que viene a decirle a la funcion que pertenece a una estructura y por ende puede acceder
a sus campos a traves de la variable myStr

*/

func main() {
	// Declarar una variable de tipo myStruct
	var myVar myStruct

	// Establecer el valor del campo FirstName a "John"
	myVar.FirstName = "John"

	// Declara una variable de tipo myStruct y
	// Establece el valor del campo FirstName a "Mary"
	myVar2 := myStruct{
		FirstName: "Mary",
	}

	// Imprime el valor de los campos FirstName de las variables
	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	// Se crea una funcion printFirstName() dentro de la estructura myStruct,
	// esta funcion regresa el valor del campo FirstName de la estructura
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
