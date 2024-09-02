package processing

import (
	"encoding/json"
	"fmt"
	"goagente/internal/data"
	"goagente/internal/logging"
	"os"
)

func CreateCoreinfoJSON() (string, error) {
	hostname, err := data.GetHostname()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o hostname CreateCoreInfoJson: %w", err)
		logging.Error(newErr)
		fmt.Println("Erro ao obter o hostname:", err)
		return "", err
	}

	username, err := data.GetCurrentUser()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o usuário atual CreateCoreInfoJson: %w", err)
		logging.Error(newErr)
		fmt.Println("Erro ao obter o usuário atual:", err)
		return "", err
	}

	// Lê o número de patrimônio do arquivo pat.txt
	patNumber, err := os.ReadFile("pat.txt")
	if err != nil {
		newErr := fmt.Errorf("erro ao ler o arquivo de patrimônio CreateCoreInfoJson: %w", err)
		logging.Error(newErr)
		return "", err
	}

	result := CoreInfoResult{
		Hostname:   hostname,
		Username:   username,
		Patrimonio: string(patNumber),
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		newErr := fmt.Errorf("erro ao converter para JSON CreateCoreInfoJson: %w", err)
		logging.Error(newErr)
		fmt.Println("Erro ao converter para JSON:", err)
		return "", err
	}

	return string(jsonBytes), nil
}

type CoreInfoResult struct {
	Hostname   string `json:"hostname"`
	Username   string `json:"username"`
	Patrimonio string `json:"patrimonio"`
}
