package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"os/exec"
)

// Estrutura para informações da placa-mãe
type MotherboardInfo struct {
	Manufacturer string `json:"Manufacturer"`
	Product      string `json:"Product"`
}

// MotherboardInfoRetriever define o contrato para obter informações da placa-mãe
type MotherboardInfoRetriever interface {
	GetMotherboardInfo() (MotherboardInfo, error)
	PowerShellGetMotherboardInfo() *exec.Cmd
	deserializeMotherboardInfo([]byte) (MotherboardInfo, error) // Método para desserializar
}

// GetMotherboardInfo retorna as informações da placa-mãe
func (m MotherboardInfo) GetMotherboardInfo() (MotherboardInfo, error) {
	cmd := m.PowerShellGetMotherboardInfo()

	// Executa o comando PowerShell
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em GetMotherboardInfo: %v", err)
		logging.Error(newErr)
		return MotherboardInfo{}, err
	}

	// Usa o novo método para desserializar o JSON
	motherboard, err := m.deserializeMotherboardInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return MotherboardInfo{}, err
	}

	return motherboard, nil
}

// PowerShellGetMotherboardInfo executa o comando PowerShell para obter informações da placa-mãe
func (MotherboardInfo) PowerShellGetMotherboardInfo() *exec.Cmd {
	// Comando PowerShell para obter informações da placa-mãe em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_BaseBoard | Select-Object -Property Manufacturer, Product | ConvertTo-Json")
	return cmd
}

// deserializeMotherboardInfo desserializa o JSON para a estrutura MotherboardInfo
func (m MotherboardInfo) deserializeMotherboardInfo(data []byte) (MotherboardInfo, error) {
	var motherboard MotherboardInfo
	err := json.Unmarshal(data, &motherboard)
	if err != nil {
		return MotherboardInfo{}, fmt.Errorf("erro ao deserializar JSON em MotherboardInfo: %v", err)
	}
	return motherboard, nil
}
