package main

import "log"

func main() {

	// En go para hacer iteraciones solo se usa el "for"
	// Para hacer un for clasico se usa la siguiente sintaxis

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	// Tambien podrias declarar primero la variable luego solo usar el for
	/*
		*
		var i2 int
		for i2 = 0; i2 < 10; i2++ {
			log.Println(i2)
		}
		*
	*/

	animales := []string{"Perro", "Gato", "Conejo", "Pez", "Tortuga"}

	// Iterar sobre un slice
	// i2 es el indice del elemento, animal es el valor de cada uno de los elementos
	// esta variable se declara y se asigna automaticamente por cada iteracion
	// Por eso se declara con :=
	for i2, animal := range animales {
		log.Println(i2, animal)
	}

	// Si no quiero el indice puedo usar el guion bajo _ para esa variable
	// NO SE PUEDE OMITIR LA VARIABLE DEL INDICE SIN MAS,
	// Ya que en vez de devolver el valor del ememento, devolveria el indice
	// Pensando que el valor no se esta usando
	for _, animal := range animales {
		log.Println(animal)
	}

	// Iterar sobre un map
	abuelos := make(map[string]string)
	abuelos["abuelaName"] = "Mary"
	abuelos["abueloName"] = "John"

	for id, value := range abuelos {
		log.Println(id, ":", value)
	}

	// Iterar sobre un string
	var str string = "Hola Mundo"

	for i, c := range str {
		log.Println(i, ":", string(c))
	}
	// Se usa el string(c) para convertir el caracter a string y poder imprimirlo
	// ya que el caracter es un tipo de dato int32, y se trabaja con la tabla ASCII

	// Crear un tipo de dato personalizado "User"
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	// Crear un slice de tipo "User" (Array Dinamico)
	var users []User

	// Agregarle 4 elementos al slice
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	// Iterar sobre el slice
	for _, user := range users {
		// Imprimir los datos de cada elemento del slice
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}
