package main

import "fmt"

const a = 10

var b string = "Hello, Go!"

type ID int

type T any

var (
	c bool
	d int
	e string
	f float64
	g byte = 'A'
	h ID   = 12345
	arr [3]int = [3]int{1, 2, 3}
)

// type System interface {
// 	DeactivateClient()
// }

// func Desativacao(client System){
// 	client.DeactivateClient()
// }

type Client struct {
	Name string
	Age int
	Active bool
	Address
}

func (client *Client) DeactivateClient() {
	client.Active = false
}

func (client Client) Deposit(value float64){
	client.Name += " - Cliente Personalite"
	fmt.Printf("client %s deposited %v\n", client.Name, value)
}

type Address struct {
	Estado string
	Cidade string
	Numero int
	Logradouro string
}

func main() {
	println("Hello, World!")
	println("The value of a is:", a)
	println("The value of b is:", b)
	b = "Welcome to Go programming!"
	println("The new value of b is:", b)
	println("The default values of c, d, e, f and g are:", c, d, e, f, g)

	shortVariable := "This is a short variable declaration."
	println(shortVariable)

	println("The value of h is:", h)

	fmt.Printf("O tipo da variável g é: %T", g)

	println("The length of the array is:", len(arr))
	println("The elements of the array are:", arr[0], arr[1], arr[2])

	for i, value := range arr {
		println("Element at index", i, "is:", arr[value-1])
	}

	// cap is the capacity of the slice, which is the total number of elements that the slice can hold. len is the length of the slice, which is the number of elements currently in the slice. When we create a slice from an array, the length of the slice is determined by the number of elements we include in the slice, while the capacity is determined by the total number of elements in the underlying array starting from the index where the slice begins. In this example, when we create a slice from index 0 to 3, it includes all three elements of the array, so both len and cap are 3. However, when we create a slice from index 0 to 0, it does not include any elements, so len is 0, but cap remains 3 because it still has access to all three elements of the underlying array.
	slice := []int{4, 5, 6, 10, -2}
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(slice[:0]), cap(slice[:0]), slice[:0]) // creates a new slice that starts at the same index as the original slice but has a length of 0. The capacity of the new slice is the same as the original slice because it still has access to all the elements of the underlying array starting from the index where the new slice begins. In this case, since the new slice starts at index 0, it has access to all 5 elements of the original slice, so its capacity is 5.
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(slice[:4]), cap(slice[:4]), slice[:4]) // creates a new slice that includes the first 4 elements of the original slice. The length of the new slice is 4 because it includes 4 elements, and the capacity is also 5 because it still has access to all the elements of the underlying array starting from the index where the new slice begins. In this case, since the new slice starts at index 0, it has access to all 5 elements of the original slice, so its capacity is 5.
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(slice[2:]), cap(slice[2:]), slice[2:]) // creates a new slice that includes all elements of the original slice starting from index 2. The length of the new slice is 3 because it includes the elements at indices 2, 3, and 4 of the original slice, and the capacity is also 3 because it only has access to the elements of the underlying array starting from index 2. In this case, since the new slice starts at index 2, it has access to the elements at indices 2, 3, and 4 of the original slice, so its capacity is 3.

	slice = append(slice, 20, 30, 60)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(slice), cap(slice), slice)

	salarys := map[string]float64{
		"John": 50000.0,
		"Jane": 60000.0,
		"Bob":  45000.0,
	}
	fmt.Printf("Salary of John: %.2f\n", salarys["John"])
	fmt.Printf("Salary of Jane: %.2f\n", salarys["Jane"])
	fmt.Printf("Salary of Bob: %.2f\n", salarys["Bob"])

	delete(salarys, "Bob")
	fmt.Printf("Salary of Bob after deletion: %.2f\n", salarys["Bob"])
	fmt.Printf("salarys map after deletion: %v\n", salarys)

	hashMap := make(map[string]int)
	hashMap["Alice"] = 25

	fmt.Printf("hasMap=%v\n", hashMap)

	for name, salary := range salarys {
		fmt.Printf("Salary: %.2f, Name: %s\n", salary, name)
	}

	result, err := sum(5, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("The sum of 5 and 10 is: %d\n", result)
	
	println("the factorial of 5 is:", fatorial(5))

	println("The result of 4 + 5 + 70 + 2.3 - 10 + 0.4444 is ", sumReduce(4, 5, 70, 2.3, -10, 0.4444))

	func() {
		println("Executado dentro do closure")
	}()

	client := Client{
		Name: "Hiago",
		Age: 32,
		Active: true,

	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", client.Name, client.Age, client.Active)
	println("Agora desativando o cliente")
	// client.Active = false;
	client.DeactivateClient()
	fmt.Printf("Ativo: %t\n", client.Active)

	client.Cidade = "Rio de Janeiro"
	client.Address.Estado = "Rio de Janeiro"
	client.Numero = 562
	client.Logradouro = "Rua 1"

	fmt.Printf("Cliente Endereco: %+v\n", client.Address)

	value := 10
	fmt.Printf("O valor de value e %d e seu endereco e o seu endereco e %v\n", value, &value)

	var pointerValue *int = &value // apontando pro local da memoria de value

	fmt.Printf("O ponteiro da variavel valor %v\n", pointerValue)

	*pointerValue = 20
	fmt.Printf("O novo ponteiro da variavel valor %v e o valor %v\n", pointerValue, value)

	otherValue := &value
	fmt.Printf("otherValue apontando pra value fica com valor de %v\n", *otherValue)

	*otherValue = 30
	fmt.Printf("otherValue: %v, value: %v\n", *otherValue, value)

	value = 50
	fmt.Printf("otherValue: %v, value: %v\n", *otherValue, value)
    // quando usar ponteiro, quando voce quer alterar os valores de entrada e nao somente utilizar os seus valores e retornar um novo
	// ou quando voce tem uma referencia do objeto 

	client.Deposit(123.56)
	println(client.Name)

	var x T = 10 // generics
	x = "Hello"

	fmt.Printf("The value of x is: %v and the type is: %T\n", x, x)

}

	func sum(a, b int) (int, error) {
		if(a < 0 || b < 0) {
			return 0, fmt.Errorf("Negative numbers are not allowed")
		}

		return a + b, nil
	}

	func fatorial(n int) int {
		if n == 0 {
			return 1
		}
		return n * fatorial(n-1)
	}

	func sumReduce(numeros ...float64) float64{ // funcao variadica - o spread do go
		total := 0.0

		for _, numero := range numeros {
			total += numero
		}

		return total
	}