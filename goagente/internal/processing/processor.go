package processing

import (
	"fmt"
	"goagente/internal/data"
	"goagente/internal/logging"
	"os"
	"regexp"
)

func CheckAndCreateFile() error {
	filePath := "pat.txt"
	// Verifica se o arquivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Se o arquivo não existe, cria-o
		file, err := os.Create(filePath)
		if err != nil {
			logging.Error(err)
			return fmt.Errorf("erro ao criar o arquivo: %v", err)
		}
		file.Close()
		fmt.Println("Arquivo criado:", filePath)
		logging.Info("Arquivo criado: " + filePath)
	}

	// Abre o arquivo para leitura
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}

	// Verifica se o arquivo está vazio
	fi, err := file.Stat()
	if err != nil {
		logging.Error(err)
		file.Close()
		return fmt.Errorf("erro ao obter informações do arquivo: %v", err)
	}

	file.Close() // Fecha o arquivo após a verificação

	if fi.Size() == 0 {
		logging.Info("O arquivo está vazio.")
		// Se o arquivo estiver vazio, abre-o para escrita
		file, err = os.OpenFile(filePath, os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("erro ao abrir o arquivo para escrita: %v", err)
		}
		defer file.Close()

		// Escreve a linha desejada
		hostname, err := data.GetHostname()
		if err != nil {
			logging.Error(err)
			return fmt.Errorf("erro ao obter o hostname: %v", err)
		}
		pat := ExtrairSequenciaFinal(hostname)

		_, err = file.WriteString(pat)
		if err != nil {
			logging.Error(err)
			return fmt.Errorf("erro ao escrever no arquivo: %v", err)
		}
		logging.Info("Linha adicionada ao arquivo patrimonio: " + pat)
		fmt.Println("Linha adicionada ao arquivo.")
	} else {
		// Se o arquivo não estiver vazio, segue o código
		fmt.Println("O arquivo não está vazio. Seguindo com o código.")
		logging.Info("O arquivo não está vazio. Seguindo com o código.")
	}

	return nil
}

func ExtrairSequenciaFinal(nomeComputador string) string {
	re := regexp.MustCompile(`[0-9]+$`)
	sequencia := re.FindString(nomeComputador)

	if sequencia == "" {
		// Retorna "00000" se nenhuma sequência numérica for encontrada
		logging.Info("Nenhuma sequência numérica encontrada no nome do computador.")
		return "00000"
	}
	return sequencia
}
