package controllers

import (
	"api/src/repositorio/cambio"
	"api/src/responses"
	"net/http"
)

func ListarMoedas(w http.ResponseWriter, r *http.Request) {
	moedas, erro := cambio.GetAll()

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, moedas)
}

func ConverterMoedas(w http.ResponseWriter, r *http.Request) {
	// Recebe o retorno das 3 variáveis
	total, totalConvertido, erro := cambio.ConverterSaldo(r)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Como não temos uma struct para esse tipo de retorno,
	// criamos uma "fantasma".
	responses.JSON(w, http.StatusOK, struct {
		Total           float64 `json:"total"`
		TotalConvertido float64 `json:"convertido"`
	}{
		Total:           total,
		TotalConvertido: totalConvertido,
	})
}

func CadastrarMoedas(w http.ResponseWriter, r *http.Request) {
	moeda, erro := cambio.CreateOne(r)

	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, moeda)
}
