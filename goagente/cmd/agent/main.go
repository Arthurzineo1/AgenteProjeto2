package main

import (
	"fmt"
	"goagente/internal/communication"
	"goagente/internal/config"
	logs "goagente/internal/logging"
	"goagente/internal/orchestration"
	"goagente/internal/processing"
	"time"
)

func main() {

	logs.Init()        // Inicializa os loggers
	defer logs.Close() // Garante que os arquivos de log serão fechados ao final da execução
	logs.Info("Aplicação iniciada com sucesso.")

	err := processing.CheckAndCreateFile()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	apiUrl := "https://run.mocky.io"
	client := communication.NewAPIClient(apiUrl)

	go orchestration.MonitorAndSendCoreInfo(client, communication.EnviaCoreInfos, config.TimeInSecondsForCoreInfoLoop)

	time.Sleep(5 * time.Second)

	orchestration.SendHardwareInfo(client, communication.EnviaHardwareInfos)

	go orchestration.SendProgramInfo(client, communication.EnviaProgramInfos, config.TimeInSecondsForProgramInfoLoop)

	select {}
}
