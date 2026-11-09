package tax

import "errors"

type TaxRepository interface {
	SaveTax(amount float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than zero")
	}
	if amount >= 1000 {
		return 10.0, nil
	}
	return 5.0, nil
}

// repositorio e uma funcao externa que nao quero testar, entao vou criar uma funcao que recebe o repositorio como parametro e testa a logica de calcular e salvar o imposto
func CalculateTaxAndSave(amount float64, repository TaxRepository) (error) {
	tax, err := CalculateTax(amount)
	if err != nil {
		return err
	}
	return repository.SaveTax(tax)
}