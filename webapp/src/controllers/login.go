package controllers

import "net/http"

// CarregarTelaDeLogin vai renderizar a tela de logn
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de login"))
}
