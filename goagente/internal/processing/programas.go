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

type ProgramInfo struct {
	Patrimonio string         `json:"patrimonio"`
	Programs   []data.Program `json:"programs"`
}

func GetProgramsInfo(client *communication.APIClient, route string) {
	// Lê o número de patrimônio do arquivo pat.txt usando os.ReadFile
	patNumber, err := os.ReadFile("pat.txt")
	if err != nil {
		logging.Error(err)
		patNumber = []byte("Patrimônio desconhecido")
	}
	for {
		programs, err := data.GetInstalledPrograms()
		if err != nil {
			logging.Error(err)
			programs = []data.Program{} // Continua com uma lista vazia de programas
		}

		programsInfo := ProgramInfo{
			Patrimonio: string(patNumber), // Adiciona o número de patrimônio
			Programs:   programs,
		}

		// Converte o programsInfo para JSON
		jsonBytes, err := json.Marshal(programsInfo)
		if err != nil {
			newErr := fmt.Errorf("Erro Marshal programas : %s", err)
			logging.Error(newErr)
		}

		// Envia o JSON para a API
		resp, err := client.GenericPost(route, jsonBytes)
		if err != nil {
			logging.Error(err)
		}
		if resp.StatusCode != 200 {
			fmt.Print("Resultado JSON:", string(jsonBytes))
			fmt.Println("Erro ao enviar as informações de programas para o servidor.")
			newErr := fmt.Errorf("erro ao enviar as informações de programas para o servidor, status: %s", resp.Status)
			logging.Error(newErr)
		} else {
			fmt.Println("Resposta do servidor:", resp.Status)
			fmt.Println("Resultado JSON:", string(jsonBytes))
			fmt.Println("Informações de programas enviadas com sucesso.")
			fmt.Println("")
			logging.Info("Informações de programas enviadas com sucesso.")
		}
		time.Sleep(30 * time.Second)
	}
}
