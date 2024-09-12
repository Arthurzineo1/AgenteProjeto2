package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"os/exec"
)

// RAM representa as informações da memória RAM
type RAM struct {
	Manufacturer string  `json:"Manufacturer"`
	Capacity     float64 `json:"Capacity"` // Capacidade em GB
	FormFactor   int     `json:"FormFactor"`
}

// RAMInfo define o contrato para obter informações da RAM
type RAMInfo interface {
	GetRAMInfo() ([]RAM, error)
	PowerShellGetRamInfo() *exec.Cmd
	deserializeRAMInfo([]byte) ([]RAM, error) // Novo método para desserializar
}

// GetRAMInfo retorna as informações da memória RAM
func (r RAM) GetRAMInfo() ([]RAM, error) {
	cmd := r.PowerShellGetRamInfo()

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

// PowerShellGetRamInfo executa o comando PowerShell para obter informações da RAM
func (RAM) PowerShellGetRamInfo() *exec.Cmd {
	// Comando PowerShell para obter informações da RAM em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_PhysicalMemory | Select-Object -Property Manufacturer, Capacity, FormFactor | ConvertTo-Json")
	return cmd
}

// deserializeRAMInfo tenta desserializar o JSON como um único objeto ou uma lista de objetos RAM
func (r RAM) deserializeRAMInfo(data []byte) ([]RAM, error) {
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
