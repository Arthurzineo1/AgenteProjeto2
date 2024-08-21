package processing

import (
	"encoding/json"
	"fmt"
	"goagente/internal/communication"
	"goagente/internal/data"
	"goagente/internal/logging"
	"os"
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
		logging.Error(err)
		return
	}
	fmt.Print("Resultado JSON:", string(jsonBytes))

	// Envia o JSON para a API
	resp, err := client.GenericPost(route, jsonBytes)
	if err != nil {
		logging.Error(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Erro ao enviar as informações para o servidor.")
		newErr := fmt.Errorf("erro ao enviar as informações para o servidor, status: %s", resp.Status)
		logging.Error(newErr)
	} else {
		fmt.Println("Informações enviadas com sucesso.")
		logging.Info("Informações de hardware enviadas com sucesso.")
	}
}
