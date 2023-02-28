package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

/*
Esta es una forma sencilla de trabajar con los datos en Go
pero hay otras formas de hacerlo que sirven para diferentes situaciones

func main() {
	var myString string
	var MyInt int
	myString = "Hello World,"
	MyInt = 42

	mySecondString := "Hello World 2,"

	log.Println(myString, mySecondString, MyInt)
}
*/

func main() {

	/*************************** MAPS ***************************/

	// Los mapas funcionan como arrays en otros lenguajes...
	// Son una colección de datos que se almacenan en pares clave-valor,
	// Funcionan muy muy rapido y son inmutales, no se pueden modificar,
	// PD: no se pueden ordenar para obtener valores por indice como en los arrays
	// array[0] array[1] => se sortena automaticamente y lo que pensabas guardabas
	// en el el indice 0, es el indice 1 o 3, etc

	myMap1 := make(map[string]string)
	myMap1["dogName"] = "Lola"
	myMap1["dogAge"] = "4"

	log.Println(myMap1["dogName"])
	log.Println(myMap1["dogAge"])

	// Un mapa que contiene un struct (User)
	myMap := make(map[string]User)

	// Crear un nuevo struct con los datos de un usuario
	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	// Guardar el struct en el mapa con la clave "me"
	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	/*************************** SLICE 1 ***************************/

	// slices -> arrays dinamicos, agregar valores a un array
	// pero sin saber cuantos valores se agregran

	// Inicio el slice vacio (el array dinamico)
	var mySlice []int

	// Le agrego 3 valores o los que quiera 2, 1, 3
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	// Ordeno el slice de menor a mayor
	sort.Ints(mySlice)

	// Imprimo el slice
	log.Println(mySlice)

	/*************************** SLICE 2 ***************************/

	// Hay otra forma de crear un slice (Array dinamico)):
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// print first two elements of slice, starting at first position
	log.Println(numbers[0:2])

	// create a slice of strings
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
