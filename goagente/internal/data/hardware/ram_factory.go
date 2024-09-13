package hardware

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

// RAM represents RAM information
type RAM struct {
	Manufacturer string  `json:"Manufacturer"`
	Capacity     float64 `json:"Capacity"` // Capacity in GB
	FormFactor   int     `json:"FormFactor"`
}

// RAMRetriever defines the contract for retrieving RAM information
type RAMRetriever interface {
	GetRAMInfo() ([]RAM, error)
}

// LinuxRAMRetriever is the Linux-specific implementation of RAMRetriever
type LinuxRAMRetriever struct{}

// NewRAMRetriever returns the correct RAMRetriever implementation based on the OS
func NewRAMRetriever() (RAMRetriever, error) {
	so := runtime.GOOS

	switch so {
	case "windows":
		return WindowsRAMRetriever{}, nil
	case "linux":
		return LinuxRAMRetriever{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", so)
	}
}

func (l LinuxRAMRetriever) GetRAMInfo() ([]RAM, error) {
	cmd := exec.Command("dmidecode", "--type", "17") // Gets RAM info using dmidecode on Linux

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error executing dmidecode: %v", err)
	}

	// Process dmidecode output
	ramList, err := l.parseDMIDecodeOutput(out.Bytes())
	if err != nil {
		return nil, err
	}

	return ramList, nil
}

// parseDMIDecodeOutput parses the output of the dmidecode command
func (l LinuxRAMRetriever) parseDMIDecodeOutput(data []byte) ([]RAM, error) {
	// Here you need to parse the dmidecode output to extract RAM info
	// This is a dummy implementation; adjust it as needed.
	var ramList []RAM
	// Parse the output and fill the ramList...
	return ramList, nil
}
