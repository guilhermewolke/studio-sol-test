package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/guilhermewolke/studio-sol-test/rest/dto"
	"github.com/guilhermewolke/studio-sol-test/rest/verifier"
)

func main() {
	http.HandleFunc("/verify", VerifyHandler)
	http.ListenAndServe(":8080", nil)
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
		return
	}

	log.Printf("VerifyHandler - Início do método")
	// Recuperando o payload enviado para o método e convertendo os dados para objetos
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var request dto.Request

	if err = json.Unmarshal(body, &request); err != nil {
		panic(err)
	}

	log.Printf("VerifyHandler - request: %#v", request)
	// Enviando para o ecossistema de validação os objetos gerados para validação, de acordo com as regras solicitadas
	response := verifier.ReleaseTheKraken(request)

	json.NewEncoder(w).Encode(response)

	// Respondendo a requisição
	log.Printf("VerifyHandler - Fim do método")
}
