package services

import (
	"encoding/json"
	"net/http"

	cep "github.com/jailtonjunior94/go-grpc/pb"
	"github.com/jailtonjunior94/go-grpc/structs"
	"golang.org/x/net/context"
)

// CepServer - Estrutura de conexão gRPC
type CepServer struct{}

// GetCep - Método de busca CEP com gRPC
func (s *CepServer) GetCep(ctx context.Context, req *cep.CepRequest) (*cep.CepResponse, error) {
	var viaCep structs.ViaCep

	response, err := http.Get("https://viacep.com.br/ws/" + req.Cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&viaCep); err != nil {
		if err != nil {
			panic(err)
		}
	}

	return &cep.CepResponse{
		Cep:         viaCep.Cep,
		Logradouro:  viaCep.Logradouro,
		Complemento: viaCep.Complemento,
		Bairro:      viaCep.Bairro,
		Localidade:  viaCep.Localidade,
		Uf:          viaCep.Uf,
		Ibge:        viaCep.Ibge,
		Gia:         viaCep.Gia,
		Ddd:         viaCep.Ddd,
		Siafi:       viaCep.Siafi,
	}, nil
}
