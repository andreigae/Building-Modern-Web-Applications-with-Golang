package main

import (
	"fmt"
	"log"
	"time"
)

// Variables globales (disponibles para todas las funciones del paquete)
var s string = "Hello World, this is a global variable"

// Definir un tipo de dato personalizado -> "User", que tendra los siguientes campos:
// FirstName, LastName, PhoneNumber, Age, BirthDate

// PD: Se podria evitar crear esta estrucrura y crear las variables de forma individual:
// var FirstName string, var LastName string, etc. Pero esto no es muy escalable
// si tenemos que pasar muchas variables a una funcion, por ejemplo, tenemos que pasar
// x num de variables, y si en el futuro tenemos que pasar (x num + 20) variables, tendriamos que cambiar
// la funcion y agregar mas variables, lo cual no es muy escalable.

// SI en cambio, creamos una estructura, podemos pasar la estructura completa
// y no tendriamos que cambiar la funcion, solo la estructura (Tipo de dato personalizado).

type User struct {
	// Los parametros estan en Mayusuclas para que sean publicos,
	// si los ponemos en minusculas, solo seran accesibles dentro del paquete (package) VER EJEMPLO 1
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	// Acceder a la variable global s, e Imprimir el valor de la variable
	log.Println(s)

	// Crear una variable de tipo User y asignarle valores a sus campos
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1 555 555 1212",
	}

	// Imprimir los valores de los campos de la variable user
	// Ver la diferencia entre usar log y fmt
	fmt.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

/*
*

EJEMPLO 1:
func ThisIsAPublicFunction() {
	Esta funcion es publica, porque esta en mayuscula
	 sera accesible desde cualquier parte del programa
}

func thisIsAPrivateFunction() {
	 Esta funcion es privada, porque esta en minuscula,
	 solo sera accesible dentro del paquete (package)
}
*
*/
