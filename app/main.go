package main

import (
	"net/http"
	"encoding/json"
	"time"
)

// Estrutura para a resposta JSON
type Response struct {
	Nome string `json:"nome"`
	Horario string `json:"horario"`
	}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	// Bloqueio de POST
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	// Resposta preenchida, buscando hora atual e convertendo para UTC e formatando para ISO 8601 (ofical da web)
	response := Response{
		Nome: "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}
	// Configura o cabeçalho para indicar que a resposta é JSON e codifica a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Endpoint para o projeto Korp
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	// Inicia o servidor na porta 8080 e garante que qualquer erro seja exibido no terminal
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}