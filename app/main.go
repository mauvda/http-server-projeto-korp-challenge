package main

import (
	"net/http"
	"encoding/json"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Estrutura para a resposta JSON
type Response struct {
	Nome string `json:"nome"`
	Horario string `json:"horario"`
	}

// Métricas do Prometheus
var (
	// Contador para o total de requisições HTTP
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP",
	},
	[]string{"path"},  //Adiciona um rótulo "path" para diferenciar as métricas por endpoint
)

	// Sinal de disponibilidade da aplicação 1 pra "Online" e 0 para "Offline"
	appAvailability = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_availability",
		Help: "Disponibilidade da aplicação (1 para disponível, 0 para indisponível)",
	},
	)
)

func init() {
	// Registra as métricas no Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(appAvailability)
	appAvailability.Set(1) // Define a disponibilidade como "Online" (1)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.WithLabelValues("/projeto-korp").Inc() // Incrementa o contador para o endpoint específico
	// Verficação da requisição GET, caso contrário retorna um erro de método não permitido
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
	// Endpoint para métricas do Prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Inicia o servidor na porta 8080 e garante que qualquer erro seja exibido no terminal
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}