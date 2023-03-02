package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/utils"

	"webapp/src/router"
)

//func init() {
//	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
//	fmt.Println(hashKey)
//
//	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
//	fmt.Println(blockKey)
//}

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
