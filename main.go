package main

import (
	"fmt"

	"github.com/marcelogbrito/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
