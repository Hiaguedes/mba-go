package main

import (
	"04-packages/client"
	"04-packages/math"
	"fmt"
)

func main() {
	println("The sum of 16 and 16 is:", math.Sum(16, 16))
	println("The package name is:", math.PACKAGE_NAME)
	// println("The module name is:", math.moduleName) // this variable is not exported, so it cannot be accessed outside the package

	client := client.NewClient("John Doe")
	client.SetDumb(true)

	fmt.Printf("Client ID: %s, Name: %s, Is Dumb: %t\n", client.GetID(), client.GetName(), client.GetIsDumb())
}