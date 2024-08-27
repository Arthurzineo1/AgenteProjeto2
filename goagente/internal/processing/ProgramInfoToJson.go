package processing

import (
	"encoding/json"
	"fmt"
	"goagente/internal/data"
	"goagente/internal/logging"
	"os"
)

type ProgramInfo struct {
	Patrimonio string         `json:"patrimonio"`
	Programs   []data.Program `json:"programs"`
}

func GetProgramsInfo() (string, error) {
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
		newErr := fmt.Errorf("erro marshal programas : %s", err)
		logging.Error(newErr)
	}
	return string(jsonBytes), nil

}
