package communication

import (
	"fmt"
	"goagente/internal/logging"
	"net/http"
)

func PostHardwareInfo(client *APIClient, route string, jsonData string) {
	resp, err := client.GenericPost(route, jsonData)
	if err != nil {
		logging.Error(err)
		fmt.Println("Erro ao enviar as informações de hardware para o servidor:", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Println("Erro ao enviar as informações de hardware para o servidor.")
		newErr := fmt.Errorf("erro ao enviar as informações de hardware para o servidor, status: %s", resp.Status)
		logging.Error(newErr)
	} else {
		fmt.Println("Resposta do servidor:", resp.Status)
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Println("Informações de hardware enviadas com sucesso.")
		fmt.Println("")
		logging.Info("Informações de hardware enviadas com sucesso.")
	}
}

func PostCoreInfo(client *APIClient, route string, jsonData string) {
	resp, err := client.GenericPost(route, jsonData)
	if err != nil {
		logging.Error(err)
		fmt.Println("Erro ao enviar as informações de core para o servidor:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Println("Erro ao enviar as informações de core para o servidor.")
		newErr := fmt.Errorf("erro ao enviar as informações de core para o servidor, rota: %s, status: %s", route, resp.Status)
		logging.Error(newErr)
	} else {
		fmt.Println("Resposta do servidor:", resp.Status)
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Println("Informações de core enviadas com sucesso.")
		fmt.Println("")
		logging.Info("Informações de core enviadas com sucesso.")
	}
}

func PostProgramInfo(client *APIClient, route string, jsonData string) {
	resp, err := client.GenericPost(route, jsonData)
	if err != nil {
		logging.Error(err)
		fmt.Println("Erro ao enviar as informações de programa para o servidor:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Println("Erro ao enviar as informações de programa para o servidor.")
		newErr := fmt.Errorf("erro ao enviar as informações de programa para o servidor, rota: %s, status: %s", route, resp.Status)
		logging.Error(newErr)
	} else {
		fmt.Println("Resposta do servidor:", resp.Status)
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Println("Informações de programa enviadas com sucesso.")
		fmt.Println("")
		logging.Info("Informações de programa enviadas com sucesso.")
	}
}
