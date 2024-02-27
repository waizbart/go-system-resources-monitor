package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Recurso representa a estrutura de dados para informações de recurso
type Recurso struct {
	CPU    float64 `json:"cpu"`
	Memoria float64 `json:"memoria"`
}

func monitorarRecursos() Recurso {
	// Obter informações da CPU
	porcentagemCPU, _ := cpu.Percent(time.Second, false)
	cpuUsage := porcentagemCPU[0]

	// Obter informações da memória
	infoMemoria, _ := mem.VirtualMemory()
	memoriaUsage := float64(infoMemoria.Used) / (1024 * 1024 * 1024) // Convertendo para GB

	return Recurso{CPU: cpuUsage, Memoria: memoriaUsage}
}

func handler(w http.ResponseWriter, r *http.Request) {
	recurso := monitorarRecursos()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recurso)
}

func main() {
	http.HandleFunc("/recursos", handler)
	porta := 8080
	fmt.Printf("Servidor em execução na porta %d...\n", porta)
	http.ListenAndServe(fmt.Sprintf(":%d", porta), nil)
}
