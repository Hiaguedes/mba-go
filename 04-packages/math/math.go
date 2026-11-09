package math

type Number interface {
	~int | ~float64
}

var moduleName = "04-packages"
var PACKAGE_NAME = "math"

func Sum[T Number](numbers ...T) T { // its a exported function because it starts with uppercase letter, and it can be used in other packages
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}
