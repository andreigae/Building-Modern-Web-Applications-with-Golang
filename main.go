package main

import (
	"log"

	"github.com/andreigae/Building-Modern-Web-Applications-with-Golang/helpers"
)

const numPool = 1000

func calculateValue(intChan chan int) {

	// Llamar a la función RandomNumber de helpers/helpers.go
	randomNumber := helpers.RandomNumber(numPool)

	// Enviar el valor al channel
	intChan <- randomNumber
}

func main() {
	// Crear un channel de tipo int
	intChan := make(chan int)

	// Defer close(intChan) para asegurarnos que el channel se cierre
	// al finalizar la función main ya que el programa puede no terminar nuca
	// y quedarse a la espera

	defer close(intChan)

	// Un ejemplo del uso de defer es el siguiente: las conexiones a la base de datos,
	// se pueden quedar abiertas para siempre si no se cierran correctamente
	// Y esto puede causar problemas de rendimiento en el servidor

	// Llamar a la función calculateValue en un goroutine (thread del sistema)
	// Se puede llamar a muchos goroutines para que se ejecuten en paralelo :)

	go calculateValue(intChan)

	// Recibir el valor del channel cuando ha finalizado la función calculateValue
	num := <-intChan
	log.Println(num)
}
