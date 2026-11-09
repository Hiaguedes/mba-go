package tax

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 {
		return 10.0
	}

	// erro em teste unitário, para simular um erro de lógica, presente em testdata
	// pra rodar o teste com o caso que deu erro so rodar  go test -run=FuzzCalculateTax/4036c7da18184c0f
	// if amount == 501.00 {
	// 	panic("Erro passou valor de 500.01")
	// }

	return 5.0
}