package hardware

// Estrutura para informações do HD
type DiskInfo struct {
	DeviceID string `json:"DeviceID"`
	Model    string `json:"Model"`
	Size     uint64 `json:"Size"` // Mantendo como uint64
}

// Estrutura para informações da placa-mãe
type MotherboardInfo struct {
	Manufacturer string `json:"Manufacturer"`
	Product      string `json:"Product"`
}

// Estrutura para informações do processador
type ProcessorInfo struct {
	Name          string `json:"Name"`
	NumberOfCores int    `json:"NumberOfCores"`
	MaxClockSpeed int    `json:"MaxClockSpeed"`
}
