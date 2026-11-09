package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	println("-----")

	for index, value := range []string{"Go", "is", "awesome"} {
		println(index, value)
	}

	println("-----")

	i := 20
	for i > 10 {
		println(i)
		i--
	}

	println("-----")

	for {
		println("This is an infinite loop")
	}

}