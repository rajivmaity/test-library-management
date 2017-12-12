package main

import (
	"fmt"
	"router"
)

func main() {

	fmt.Println("Welcome to the server")

	echoInstancePointer := router.New()

	echoInstancePointer.Start(":8081")

}
