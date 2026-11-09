package main

type Number interface { // constraint that allows int and float64 types
	~int | ~float64 // the ~ operator allows any type that has int or float64 as underlying type
}

func sum[T Number](m map[string]T) T {
	var result T

	for _, value := range m {
		result += value
	}

	return result
}

func Compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	salarys := map[string]float64{
		"Mary": 1000.0,
		"John": 2500.0,
		"Jane": 3500.60,
	}

	sum := sum(salarys)

	println("The sum of the salarys is:", sum)

	println("Comparing 10 and 10.0:", Compare(10, 10.0))
	println("Comparing 10 and 10.0:", Compare("a", "A"))

}