package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[2:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching data for CEP %s: %v\n", cep, err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response for CEP %s: %v\n", cep, err)
		}

		var cepResponse CEPResponse

		err = json.Unmarshal(res, &cepResponse)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON for CEP %s: %v\n", cep, err)
		}

		fmt.Printf("response for CEP: %+v\n", cepResponse)

		// ceps registados em um arquivo de texto, um por linha, podem ser lidos usando os pacotes "os" e "bufio". O código a seguir demonstra como escrever um arquivo de texto contendo todos os registros de CEPs, um por linha, e fazer uma requisição para cada CEP usando a API do ViaCEP:
		file, err := os.OpenFile("11-busca-CEP/ceps.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		file.WriteString(string(res))

		defer file.Close()


	}
}