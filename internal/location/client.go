package location

import (
	"encoding/json"
	"fmt"
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
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

type LocationClient struct{}

func NewLocationClient() *LocationClient {
	return &LocationClient{}
}

func (lc *LocationClient) GetLocation(cep string) (*ViaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch location for CEP: %s", cep)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var viaCEP ViaCEP
	err = json.Unmarshal(body, &viaCEP)
	if err != nil {
		return nil, err
	}

	if viaCEP.Localidade == "" {
		return nil, fmt.Errorf("can not find zipcode")
	}

	return &viaCEP, nil
}
