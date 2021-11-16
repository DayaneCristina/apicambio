package rotas

import (
	"api/src/controllers"
	"net/http"
)

var cambio = []Rota{
	{
		URI:    "/moedas",
		Metodo: http.MethodGet,
		Funcao: controllers.ListarMoedas,

		RequerAutenticacao: false,
	},
	{
		URI:    "/converter/{moeda}",
		Metodo: http.MethodGet,
		Funcao: controllers.ConverterMoedas,

		RequerAutenticacao: false,
	},
	{
		URI:    "/cadastrar-moeda",
		Metodo: http.MethodPost,
		Funcao: controllers.CadastrarMoedas,

		RequerAutenticacao: false,
	},
}
