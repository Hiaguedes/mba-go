package main

import (
	"encoding/json"
	"net/http"
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
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/cep", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain") // Header is used to set the HTTP headers of the response. In this case, it sets the Content-Type header to "text/plain" to indicate that the response body contains plain text.
	w.WriteHeader(http.StatusOK) // WriteHeader is used to set the HTTP status code of the response. In this case, it sets the status code to 200 OK to indicate that the request was successful.
	w.Write([]byte("Hello, World!"))
}

func SearchCEP(cep string) (*CEPResponse, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/") // Get is used to make an HTTP GET request to the specified URL. It takes the URL as an argument and returns the response and an error if any.

	if err != nil {
		return nil, err
	}

	var cepResponse CEPResponse

	err = json.NewDecoder(req.Body).Decode(&cepResponse)

	if err != nil {
		return nil, err
	}

	defer req.Body.Close() 
	return &cepResponse, nil
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cep" {
		w.WriteHeader(http.StatusNotFound) // WriteHeader is used to set the HTTP status code of the response. In this case, it sets the status code to 404 Not Found if the requested URL path is not "/".
		return
	}
	cepParam := r.URL.Query().Get("cep") // Query is used to access the query parameters of the URL. In this case, it retrieves the value of the "cep" query parameter.
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest) // WriteHeader is used to set the HTTP status code of the response. In this case, it sets the status code to 400 Bad Request if the "cep" query parameter is not provided.
		w.Write([]byte("Missing 'cep' query parameter"))
		return
	}

	cepResponse, err := SearchCEP(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // WriteHeader is used to set the HTTP status code of the response. In this case, it sets the status code to 500 Internal Server Error if there's an error searching for the CEP.
		w.Write([]byte("Error searching CEP: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json") // Header is used to set the HTTP headers of the response. In this case, it sets the Content-Type header to "application/json" to indicate that the response body contains JSON data.
	w.WriteHeader(http.StatusOK) // WriteHeader is used to set the HTTP status code of the response. In this case, it sets the status code to 200 OK to indicate that the request was successful.
	json.NewEncoder(w).Encode(cepResponse)
}
