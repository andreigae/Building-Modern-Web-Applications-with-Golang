package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Defino una estructura que se corresponde con el json que quiero deserializar :)
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	// Mi json que quiero deserializar, que se recibirá de un API etc.
	myJson := `
	[
		{
			"first_name": "Andrei",
			"last_name": "Gae",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	// Crear una varialbe llamada unmarshalled de tipo slice de Person (Array dinámico)
	var unmarshalled []Person

	// Usar la funcion Unmarshall de Json para deserializar,
	// pasándole el json como "myJson" y la variable donde se guardará el resultado
	// Se la pasa el puntero de la variable a la funcion de Unmarshall ya que
	// Esta variable se va a modificar dentro de la función externa

	// El resultado de esta funcion se guarda en la variable err
	// Para comprobar si ha habido algún error en la deserialización
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	// Si ha habido algún error, se imprime por pantalla que ha habido un error
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	// Si no ha habido ningún error, se imprime por pantalla
	// el resultado de la deserialización
	log.Printf("unmarshalled: %v", unmarshalled)

	for _, person := range unmarshalled {
		fmt.Println("First Name: ", person.FirstName)
		fmt.Println("Last Name: ", person.LastName)
		fmt.Println("Hair Color: ", person.HairColor)
		fmt.Println("Has Dog: ", person.HasDog)
		fmt.Println()
	}

	fmt.Println("-------------------------------------------------")
	fmt.Println()

	// Escritura de un json a partir de una estructura

	// Declarar un slice de Person (Array dinámico)
	var mySlice []Person

	// Crear una estructura de tipo Person
	var m1 Person
	m1.FirstName = "Persona 1 Frist Name"
	m1.LastName = "Persona 1 Last Name"
	m1.HairColor = "Persona 1 Hair color"
	m1.HasDog = false

	// Guardar en el slice la estructura m1
	mySlice = append(mySlice, m1)

	// Crear 2º estructura de tipo Person
	var m2 Person
	m2.FirstName = "Persona 2 Frist Name"
	m2.LastName = "Persona 2 Last Name"
	m2.HairColor = "Persona 2 Hair color"
	m2.HasDog = false

	// Guardar en el slice la estructura m2
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
