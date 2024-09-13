package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"os/exec"
)

// WindowsRAMRetriever implementa o método GetRAMInfo para Windows
type WindowsRAMRetriever struct{}

// GetRAMInfo retorna as informações da memória RAM no Windows
func (r WindowsRAMRetriever) GetRAMInfo() ([]RAM, error) {
	cmd := r.powerShellGetRamInfo()

	// Executa o comando PowerShell
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell: %v", err)
		logging.Error(newErr)
		return nil, newErr
	}

	// Usa o novo método para desserializar o JSON
	ramList, err := r.deserializeRAMInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	// Converte a capacidade de cada RAM para gigabytes
	for i := range ramList {
		ramList[i].Capacity = utils.BytesToGigabytes(uint64(ramList[i].Capacity))
	}

	return ramList, nil
}

// powerShellGetRamInfo executa o comando PowerShell para obter informações da RAM
func (WindowsRAMRetriever) powerShellGetRamInfo() *exec.Cmd {
	return exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_PhysicalMemory | Select-Object -Property Manufacturer, Capacity, FormFactor | ConvertTo-Json")
}

// deserializeRAMInfo desserializa o JSON para informações da RAM
func (WindowsRAMRetriever) deserializeRAMInfo(data []byte) ([]RAM, error) {
	// Tenta deserializar o JSON como um único objeto
	var singleRAM RAM
	err := json.Unmarshal(data, &singleRAM)
	if err == nil {
		// Sucesso, retorna como uma lista de um único item
		return []RAM{singleRAM}, nil
	}

	// Se falhar, tenta deserializar como um array de objetos
	var ramList []RAM
	err = json.Unmarshal(data, &ramList)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar JSON em RAMInfo: %v", err)
	}

	return ramList, nil
}
