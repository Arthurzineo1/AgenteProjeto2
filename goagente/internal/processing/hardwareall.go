package processing

import (
	"encoding/json"
	"fmt"
	"goagente/internal/communication"
	"goagente/internal/data"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"math"
	"os"
)

func GetHardwareInfo(client *communication.APIClient, route string) {
	var jsonResult string

	// Lê o número de patrimônio do arquivo pat.txt usando os.ReadFile
	patNumber, err := os.ReadFile("pat.txt")
	if err != nil {
		logging.Error(err)
		patNumber = []byte("Patrimônio desconhecido")
	}

	disks, err := data.GetDiskInfo()
	if err != nil {
		logging.Error(err)
		disks = []data.DiskInfo{} // Continua com uma lista vazia de discos
	}
	for i := range disks {
		disks[i].Size = uint64(math.Round(utils.BytesToGigabytes(disks[i].Size)))
	}

	processors, err := data.GetProcessorInfo()
	if err != nil {
		logging.Error(err)
		processors = []data.ProcessorInfo{} // Continua com uma lista vazia de processadores
	}

	ram, err := data.GetRAMInfo()
	if err != nil {
		logging.Error(err)
		ram = []data.RAMInfo{} // Continua com uma lista vazia de módulos RAM
	}

	motherboard, err := data.GetMotherboardInfo()
	if err != nil {
		logging.Error(err)
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
		logging.Error(err)
	} else {
		jsonResult = string(jsonBytes)
		logging.Info("Informações de hardware obtidas com sucesso")
		logging.Info("Resultado JSON: " + jsonResult)
	}

	resp, err := client.GenericPost(route, jsonResult)
	if err != nil {
		logging.Error(err)
		fmt.Println("Erro ao enviar as informações para o servidor:", err)
		return
	}
	fmt.Println("Resultado JSON:", jsonResult)
	fmt.Println("Resposta do servidor:", resp.Status)
	if resp.StatusCode != 200 {
		fmt.Println("Erro ao enviar as informações para o servidor.")
		newErr := fmt.Errorf("erro ao enviar as informações para o servidor, status: %s", resp.Status)
		logging.Error(newErr)
	} else {
		fmt.Println("Informações enviadas com sucesso.")
		logging.Info("Informações de hardware enviadas com sucesso.")
	}

	// Se você quiser enviar o JSON para algum lugar, faça isso aqui
	// client.GenericPost(url, jsonResult) - exemplo de como poderia ser usado
}

type HardwareInfo struct {
	Patrimonio  string               `json:"patrimonio"`
	Disks       []data.DiskInfo      `json:"disks"`
	Processors  []data.ProcessorInfo `json:"processors"`
	RAMModules  []data.RAMInfo       `json:"ram"`
	Motherboard data.MotherboardInfo `json:"motherboard"`
}
