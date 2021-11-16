package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota do API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar Rota
func Configurar(r *mux.Router) *mux.Router {
	rotas := operacao
	rotas = append(rotas, cambio...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
