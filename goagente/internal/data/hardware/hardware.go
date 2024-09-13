package hardware

// HardwareInfoFactory define os métodos para obter informações de hardware
type HardwareInfoFactory interface {
	GetDiskInfo() ([]DiskInfo, error)
	GetMotherboardInfo() (MotherboardInfo, error)
	GetProcessorInfo() ([]ProcessorInfo, error)
	GetRAMInfo() ([]RAM, error)
}
