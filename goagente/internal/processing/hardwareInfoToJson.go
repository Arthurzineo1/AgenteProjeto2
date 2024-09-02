package processing

import (
	"encoding/json"
	"fmt"
	"goagente/internal/data"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"math"
	"os"
)

func CreateHardwareInfoJSON() (string, error) {
	// Lê o número de patrimônio do arquivo pat.txt usando os.ReadFile
	patNumber, err := os.ReadFile("pat.txt")
	if err != nil {
		newErr := fmt.Errorf("erro ao ler o arquivo de patrimônio CreateHardwareInfoJSON: %w", err)
		logging.Error(newErr)
		patNumber = []byte("Patrimônio desconhecido")
		return "", err

	}

	disks, err := data.GetDiskInfo()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter informações do disco CreateHardwareInfoJSON: %w", err)
		logging.Error(newErr)
		disks = []data.DiskInfo{} // Continua com uma lista vazia de discos

	}
	for i := range disks {
		disks[i].Size = uint64(math.Round(utils.BytesToGigabytes(disks[i].Size)))
	}

	processors, err := data.GetProcessorInfo()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter informações do processador CreateHardwareInfoJSON: %w", err)
		logging.Error(newErr)
		processors = []data.ProcessorInfo{} // Continua com uma lista vazia de processadores
	}

	ram, err := data.GetRAMInfo()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter informações da RAM CreateHardwareInfoJSON: %w", err)
		logging.Error(newErr)
		ram = []data.RAMInfo{} // Continua com uma lista vazia de módulos RAM
	}

	motherboard, err := data.GetMotherboardInfo()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter informações da placa-mãe CreateHardwareInfoJSON: %w", err)
		logging.Error(newErr)
		motherboard = data.MotherboardInfo{} // Continua com uma estrutura vazia de placa-mãe
	}

	hardwareInfo := HardwareInfo{
		Patrimonio:  string(patNumber), // Adiciona o número de patrimônio como o primeiro campo
		Disks:       disks,
		Processors:  processors,
		RAMModules:  ram,
		Motherboard: motherboard,
	}

	// Converte o hardwareInfo para JSON
	jsonBytes, err := json.Marshal(hardwareInfo)
	if err != nil {
		newErr := fmt.Errorf("erro ao converter para JSON CreateHardwareInfoJSON: %w", err)
		logging.Error(newErr)
		return "", err
	}

	return string(jsonBytes), nil
}

type HardwareInfo struct {
	Patrimonio  string               `json:"patrimonio"`
	Disks       []data.DiskInfo      `json:"disks"`
	Processors  []data.ProcessorInfo `json:"processors"`
	RAMModules  []data.RAMInfo       `json:"ram"`
	Motherboard data.MotherboardInfo `json:"motherboard"`
}
