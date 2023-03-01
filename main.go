package main

import (
	"fmt"
	"log"

	"github.com/andreigae/Building-Modern-Web-Applications-with-Golang/helpers"
)

func main() {
	var user helpers.User
	user.FirstName = "Andrei"
	user.LastName = "Gae"
	fmt.Println(user.FirstName, user.LastName)
	log.Println(user.FirstName, user.LastName)
}
