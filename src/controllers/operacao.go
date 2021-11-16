package controllers

import (
	deposito "api/src/repositorio/deposito"
	"api/src/responses"
	"net/http"
)

func Depositar(w http.ResponseWriter, r *http.Request) {
	deposito, erro := deposito.CreateOne(r)

	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)

		return
	}

	responses.JSON(w, http.StatusOK, deposito)
}

func ListarTodos(w http.ResponseWriter, r *http.Request) {
	depositos, erro := deposito.GetAll()

	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return

	}

	responses.JSON(w, http.StatusAccepted, depositos)

}
