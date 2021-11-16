package controllers

import "net/http"

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando"))
}
