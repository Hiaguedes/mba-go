package main

import "fmt"

type Number interface {
	~int | ~float64
}

type StatusType interface {
	~string
}

const (
	ERROR    = "ERROR"
	SUCCESS  = "SUCCESS"
)

func Compare[T Number](a, b T) {
	if a > b {
		fmt.Printf("%d is greater than %d\n", a, b)
	} else if a < b {
		fmt.Printf("%d is less than %d\n", a, b)
	} else {
		fmt.Printf("%d is equal to %d\n", a, b)
	}
}

func Status[T StatusType](status T) {
	switch status {
	case ERROR:
		fmt.Println("An error occurred")
	case SUCCESS:
		fmt.Println("Operation was successful")
	default:
		fmt.Println("Unknown status")
	}
}

func main() {
	Compare(10, 20)
	Compare(20, 10)
	Compare(10, 10)

	Status(ERROR)
	Status(SUCCESS)
	Status("UNKNOWN")

}