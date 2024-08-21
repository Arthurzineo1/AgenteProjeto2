package data

import (
	"golang.org/x/sys/windows/registry"
)

type Program struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// GetInstalledPrograms retorna uma lista de programas instalados.
func GetInstalledPrograms() ([]Program, error) {
	var programs []Program

	// Caminho do registro onde os programas instalados estão listados
	keys := []string{
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}

	for _, key := range keys {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.READ)
		if err != nil {
			continue
		}
		defer k.Close()

		names, err := k.ReadSubKeyNames(-1)
		if err != nil {
			continue
		}

		for _, name := range names {
			subKey, err := registry.OpenKey(k, name, registry.READ)
			if err != nil {
				continue
			}
			defer subKey.Close()

			displayName, _, err := subKey.GetStringValue("DisplayName")
			if err != nil {
				continue
			}

			displayVersion, _, err := subKey.GetStringValue("DisplayVersion")
			if err != nil {
				displayVersion = "N/A"
			}

			programs = append(programs, Program{Name: displayName, Version: displayVersion})
		}
	}

	return programs, nil
}
