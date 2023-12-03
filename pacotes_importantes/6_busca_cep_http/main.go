package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCEPHandler)

	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BuscaCEP(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, error := json.Marshal(cep)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func BuscaCEP(cep string) (*ViaCEP, error) {
	res, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer res.Body.Close()

	body, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCEP
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil
}
