package hardware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"goagente/pkg/utils"
	"os/exec"
)

// Estrutura para informações do HD
type DiskInfo struct {
	DeviceID string `json:"DeviceID"`
	Model    string `json:"Model"`
	Size     uint64 `json:"Size"` // Mantendo como uint64
}

// DiskInfoRetriever define o contrato para obter informações do disco
type DiskInfoRetriever interface {
	GetDiskInfo() ([]DiskInfo, error)
	PowerShellGetDiskInfo() *exec.Cmd
	deserializeDiskInfo([]byte) ([]DiskInfo, error) // Método para desserializar
}

// GetDiskInfo retorna as informações do disco
func (d DiskInfo) GetDiskInfo() ([]DiskInfo, error) {
	cmd := d.PowerShellGetDiskInfo()

	// Executa o comando PowerShell
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando PowerShell em getdiskinfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Usa o novo método para desserializar o JSON
	disks, err := d.deserializeDiskInfo(out.Bytes())
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	// Converte o tamanho para gigabytes
	for i := range disks {
		disks[i].Size = uint64(utils.BytesToGigabytes(disks[i].Size))
	}

	return disks, nil
}

// PowerShellGetDiskInfo executa o comando PowerShell para obter informações do disco
func (DiskInfo) PowerShellGetDiskInfo() *exec.Cmd {
	// Comando PowerShell para obter informações do disco em formato JSON
	cmd := exec.Command("powershell", "-Command", "Get-PhysicalDisk | Select-Object -Property DeviceID, Model, Size | ConvertTo-Json")
	return cmd
}

// deserializeDiskInfo tenta desserializar o JSON como um único objeto ou uma lista de objetos DiskInfo
func (d DiskInfo) deserializeDiskInfo(data []byte) ([]DiskInfo, error) {
	// Tenta deserializar o JSON como um único objeto
	var singleDisk DiskInfo
	err := json.Unmarshal(data, &singleDisk)
	if err == nil {
		// Se for um único objeto, retorna como um slice com um único item
		return []DiskInfo{singleDisk}, nil
	}

	// Se falhar, tenta deserializar como um array de objetos
	var disks []DiskInfo
	err = json.Unmarshal(data, &disks)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar JSON em diskInfo: %v", err)
	}

	return disks, nil
}
