package tax

import (
	"testing"
)

// comandos pra rodar testes

// go test .
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out

func TestCalculateTax(t *testing.T) {
	// arrange
	amount := 500.0

	// act
	tax := CalculateTax(amount)

	// assert
	if tax != 5.0 {
		t.Errorf("Expected tax to be 5.0, but got %f", tax)
	}

	// arrange 
	amount = 1500.0
	expectedTax := 10.0

	// act
	tax = CalculateTax(amount)

	// assert
	if tax != expectedTax {
		t.Errorf("Expected tax to be %f, but got %f", expectedTax, tax)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount       float64
		expectedTax  float64
	}

	testCases := []calcTax{
		{amount: 500.0, expectedTax: 5.0},
		{amount: 1500.0, expectedTax: 10.0},
		{amount: 1000.0, expectedTax: 10.0},
		{amount: 999.99, expectedTax: 5.0},
		{amount: 0.0, expectedTax: 0.0},
		{amount: -100.0, expectedTax: 0.0},
	}

	for _, tc := range testCases {
		t.Run("CalculateTax", func(t *testing.T) {
			tax := CalculateTax(tc.amount)
			if tax != tc.expectedTax {
				t.Errorf("Expected tax to be %f for amount %f, but got %f", tc.expectedTax, tc.amount, tax)
			}
		})
	}
}

// um loop que roda o teste várias vezes, para medir o tempo de execução
// rodo com go test -bench=.
// pra ter uma contagem e executar o benchmark n vezes eu passo go test -bench=. -run=^# -count=n (onde n pode ser 10, 100, 1000, etc, qq numero inteiro positivo)
// pra conseguir fazer benchmark vendo quanto de memoria foi alocada, eu rodo com go test -bench=. -benchmem

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

// pra testes de mutacao, eu passo go test -fuzz=Fuzz -fuzztime=10s
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{500.0, 1500.0, 1000.0, 999.99, 0.0, -100.0, -1.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		tax := CalculateTax(amount)

		if amount <= 0 && tax != 0.0 {
			t.Errorf("Expected tax to be 0.0 for amount %f, but got %f", amount, tax)
		}

		if amount > 0 && amount < 1000 && tax != 5.0 {
			t.Errorf("Expected tax to be 5.0 for amount %f, but got %f", amount, tax)
		}

		if amount >= 1000 && tax != 10.0 {
			t.Errorf("Expected tax to be 10.0 for amount %f, but got %f", amount, tax)
		}
	})
}