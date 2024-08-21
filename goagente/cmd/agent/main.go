package main

import (
	"fmt"
	"goagente/internal/communication"
	logs "goagente/internal/logging"
	"goagente/internal/processing"
)

func main() {

	logs.Init()        // Inicializa os loggers
	defer logs.Close() // Garante que os arquivos de log serão fechados ao final da execução
	logs.Info("Aplicação iniciada com sucesso.")
	// Verifica se o arquivo com o pat existe e o cria, se necessário
	err := processing.CheckAndCreateFile()
	if err != nil {
		fmt.Println("Erro:", err)
	}
	apiUrl := "https://run.mocky.io"
	client := communication.NewAPIClient(apiUrl)

	go processing.CoreInfos(client, communication.EnviaCoreInfos)           // executado em uma goroutine looping infinito com sleep de 10 segundos (vai ser aumentado)
	go processing.GetHardwareInfo(client, communication.EnviaHardwareInfos) // executado apenas 1 vez quando o agente é
	go processing.GetProgramsInfo(client, communication.EnviaProgramInfos)  // executado apenas 1 vez quando o agente é iniciado

	select {}
}
