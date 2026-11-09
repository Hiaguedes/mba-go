package sum // Sum returns the sum of two integers. the name of the package is sum, and the name of the function is Sum. The function takes two integers as parameters and returns their sum as an integer.
// the name of the package here must be the same as the name of the folder where this file is located, which is sum. This is a convention in Go, and it allows us to import this package in other files using the package name. For example, if we want to use the Sum function in another file, we can import this package and call the function like this: sum.Sum(2, 3).

func Sum(a, b int) int {
	return a + b
}