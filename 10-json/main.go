package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  float64 `json:"s"`
}

func main() {
	conta := Conta{Numero: 12345, Saldo: 1000.50} // Create an instance of the Conta struct with the specified values for Numero and Saldo.


	jsonData, err := json.Marshal(conta) // Marshal is used to convert the Conta struct into JSON format. It takes the struct as an argument and returns the JSON data as a byte slice and an error if any.

	if err != nil {
		panic(err)
	}

	println("JSON data:", string(jsonData)) // Print the JSON data. The byte slice is converted to a string for printing.
	encoder := json.NewEncoder(os.Stdout) // NewEncoder is used to create a new JSON encoder that writes to the standard output. It takes the output destination as an argument and returns a JSON encoder.

	err = encoder.Encode(conta) // Encode is used to encode the Conta struct into JSON format and write it to the standard output. It takes the struct as an argument and returns an error if any.

	if err != nil {
		panic(err)
	}

	var contaX Conta // Create a variable of type Conta to hold the decoded data.

	err = json.Unmarshal(jsonData, &contaX) // Unmarshal is used to convert the JSON data back into a Conta struct. It takes the JSON data and a pointer to the struct as arguments and returns an error if any.

	if err != nil {
		panic(err)
	}

	fmt.Printf("Conta decodificada: %+v\n", contaX)

	pureJson := []byte(`{"n":54321,"s":2000.75}`) // Create a string containing JSON data that represents a Conta struct with the specified values for Numero and Saldo.

	var contaY Conta // Create a variable of type Conta to hold the decoded data.

	err = json.Unmarshal([]byte(pureJson), &contaY) // Unmarshal is used to convert the JSON data in the pureJson string back into a Conta struct. It takes the JSON data as a byte slice and a pointer to the struct as arguments and returns an error if any.

	if err != nil {
		panic(err)
	}

	fmt.Printf("Conta decodificada a partir de string JSON: %+v\n", contaY) // Print the decoded Conta struct. The %+v verb is used to print the struct with field names.

	
}