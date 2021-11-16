package rotas

import (
	"api/src/controllers"
	"net/http"
)

var operacao = []Rota{
	{
		URI:                "/depositar",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Depositar,
		RequerAutenticacao: false,
	},

	{
		URI:                "/listar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarTodos,
		RequerAutenticacao: false,
	},
}
