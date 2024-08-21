package processing

import (
	"encoding/json"
	"fmt"
	"goagente/internal/communication"
	"goagente/internal/data"
	"goagente/internal/logging"
	"os"
	"time"
)

func CoreInfos(client *communication.APIClient, route string) {
	var jsonResult string
	for {
		hostname, err := data.GetHostname()
		if err != nil {
			logging.Error(err)
			fmt.Println("erro ao obter o hostname:", err)
			continue
		}

		username, err := data.GetCurrentUser()
		if err != nil {
			logging.Error(err)
			fmt.Println("erro ao obter o usuário atual:", err)
			continue
		}

		// Lê o número de patrimônio do arquivo pat.txt
		patNumber, err := os.ReadFile("pat.txt")
		if err != nil {
			logging.Error(fmt.Errorf("erro ao ler o arquivo de patrimônio: %w", err))
			patNumber = []byte("patrimônio desconhecido")
		}

		result := CoreInfoResult{
			Hostname:   hostname,
			Username:   username,
			Patrimonio: string(patNumber),
			Timestamp:  time.Now().Format(time.RFC3339),
		}

		jsonBytes, err := json.Marshal(result)
		if err != nil {
			logging.Error(err)
			fmt.Println("erro ao converter para JSON:", err)
			continue
		}
		jsonResult = string(jsonBytes)
		resp, err := client.GenericPost(route, jsonResult)
		if err != nil {
			logging.Error(err)
			fmt.Println("erro ao enviar as informações Core para o servidor:", err)
			continue
		}
		fmt.Println("Resposta do servidor:", resp.Status)
		if resp.StatusCode != 200 {
			fmt.Println("Resultado JSON:", jsonResult)
			fmt.Println("erro ao enviar as informações Core para o servidor.")
			newErr := fmt.Errorf("erro ao enviar as informações Core para o servidor, status: %s", resp.Status)
			logging.Error(newErr)
		} else {
			fmt.Println("Resultado JSON:", jsonResult)
			fmt.Println("Informações Core enviadas com sucesso.")
			fmt.Println("")
			logging.Info("informações Core enviadas com sucesso.")
		}

		time.Sleep(10 * time.Second)
	}
}

type CoreInfoResult struct {
	Hostname   string `json:"hostname"`
	Username   string `json:"username"`
	Patrimonio string `json:"patrimonio"`
	Timestamp  string `json:"timestamp"`
}
