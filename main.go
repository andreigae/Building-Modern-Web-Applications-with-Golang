package main

import (
	"github.com/andreigae/Building-Modern-Web-Applications-with-Golang/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
