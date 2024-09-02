package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goagente/internal/logging"
	"os/exec"
)

// Estrutura para informações do HD
type DiskInfo struct {
	DeviceID string `json:"DeviceID"`
	Model    string `json:"Model"`
	Size     uint64 `json:"Size"` // Mantendo como uint64
}

// Estrutura para informações do processador
type ProcessorInfo struct {
	Name          string `json:"Name"`
	NumberOfCores int    `json:"NumberOfCores"`
	MaxClockSpeed int    `json:"MaxClockSpeed"`
}

// Estrutura para informações da RAM
type RAMInfo struct {
	Manufacturer string  `json:"Manufacturer"`
	Capacity     float64 `json:"Capacity"` // Alterado para armazenar a capacidade em GB
	FormFactor   int     `json:"FormFactor"`
}

// Estrutura para informações da placa-mãe
type MotherboardInfo struct {
	Manufacturer string `json:"Manufacturer"`
	Product      string `json:"Product"`
}

// Função para converter bytes em gigabytes
func BytesToGigabytes(bytes uint64) float64 {
	const bytesInGigabyte = 1024 * 1024 * 1024
	return float64(bytes) / bytesInGigabyte
}

func GetDiskInfo() ([]DiskInfo, error) {
	cmd := exec.Command("powershell", "-Command", "Get-PhysicalDisk | Select-Object -Property DeviceID, Model, Size | ConvertTo-Json")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando powershell em getdiskinfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Verifica e imprime o JSON retornado para debugging
	//logging.Info(out.String())

	// A primeira tentativa de deserialização para um único objeto DiskInfo
	var singleDisk DiskInfo
	err = json.Unmarshal(out.Bytes(), &singleDisk)
	if err == nil {

		// Se for um único objeto, coloca em um slice
		return []DiskInfo{singleDisk}, nil

	}

	// Se falhar como objeto único, tenta deserializar como uma lista de objetos
	var disks []DiskInfo
	err = json.Unmarshal(out.Bytes(), &disks)
	if err != nil {
		newErr := fmt.Errorf("erro ao deserializar JSON em diskInfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Converte o tamanho para gigabytes
	for i := range disks {
		disks[i].Size = uint64(BytesToGigabytes(disks[i].Size))
	}

	return disks, nil
}

func GetProcessorInfo() ([]ProcessorInfo, error) {
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_Processor | Select-Object -Property Name, NumberOfCores, MaxClockSpeed | ConvertTo-Json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando powershell em GetProcessorInfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Verifica se a saída JSON é um objeto ou uma lista de objetos
	var singleProcessor ProcessorInfo
	var processors []ProcessorInfo

	// Tenta deserializar como um objeto único
	if err = json.Unmarshal(out.Bytes(), &singleProcessor); err == nil {
		// Se bem-sucedido, adiciona o objeto à lista de processadores
		processors = append(processors, singleProcessor)
	} else {
		// Se falhar, tenta deserializar como uma lista de objetos
		if err = json.Unmarshal(out.Bytes(), &processors); err != nil {
			return nil, err
		}
	}

	return processors, nil
}
func GetRAMInfo() ([]RAMInfo, error) {
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_PhysicalMemory | Select-Object -Property Manufacturer, Capacity, FormFactor | ConvertTo-Json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando powershell em GetRAMInfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// A primeira tentativa de deserialização para um único objeto RAMInfo
	var singleRAM RAMInfo
	err = json.Unmarshal(out.Bytes(), &singleRAM)
	if err == nil {
		// Se for um único objeto, coloca em um slice
		singleRAM.Capacity = BytesToGigabytes(uint64(singleRAM.Capacity))
		return []RAMInfo{singleRAM}, nil
	}

	// Se falhar como objeto único, tenta deserializar como uma lista de objetos
	var ram []RAMInfo
	err = json.Unmarshal(out.Bytes(), &ram)
	if err != nil {
		newErr := fmt.Errorf("erro ao deserializar JSON em RAMInfo: %v", err)
		logging.Error(newErr)
		return nil, err
	}

	// Converte a capacidade para gigabytes
	for i := range ram {
		ram[i].Capacity = BytesToGigabytes(uint64(ram[i].Capacity))
	}

	return ram, nil
}

func GetMotherboardInfo() (MotherboardInfo, error) {
	cmd := exec.Command("powershell", "-Command", "Get-WmiObject -Class Win32_BaseBoard | Select-Object -Property Manufacturer, Product | ConvertTo-Json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		newErr := fmt.Errorf("erro ao executar o comando powershell em GetMotherboardInfo: %v", err)
		logging.Error(newErr)
		return MotherboardInfo{}, err
	}

	var motherboard MotherboardInfo
	err = json.Unmarshal(out.Bytes(), &motherboard)
	if err != nil {
		newErr := fmt.Errorf("erro ao deserializar JSON em MotherboardInfo: %v", err)
		logging.Error(newErr)
		return MotherboardInfo{}, err
	}

	return motherboard, nil
}
