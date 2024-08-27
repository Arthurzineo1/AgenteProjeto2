package orchestration

import (
	"goagente/internal/communication"
	"goagente/internal/logging"
	"goagente/internal/processing"
	"time"
)

// SendHardwareInfo collects hardware information and sends it using the communication layer.
func SendHardwareInfo(client *communication.APIClient, route string) {
	jsonHardware, err := processing.CreateHardwareInfoJSON()
	if err != nil {
		// Log the error and return to avoid sending incomplete data
		logging.Error(err)
		return
	}
	// Send the hardware information
	communication.PostHardwareInfo(client, route, jsonHardware)
}

// MonitorAndSendCoreInfo continuously monitors and sends core information at specified intervals.
func MonitorAndSendCoreInfo(client *communication.APIClient, route string, seconds int) {
	for {
		jsonCore, err := processing.CreateCoreinfoJSON()
		if err != nil {
			// Log the error and continue the loop
			logging.Error(err)
			continue
		}
		// Send the core information
		communication.PostCoreInfo(client, route, jsonCore)

		// Wait for the specified interval before sending again
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

func SendProgramInfo(client *communication.APIClient, route string, seconds int) {
	for {
		jsonProgram, err := processing.GetProgramsInfo()
		if err != nil {
			// Log the error and return to avoid sending incomplete data
			logging.Error(err)
			return
		}
		// Send the hardware information
		communication.PostHardwareInfo(client, route, jsonProgram)

		time.Sleep(time.Duration(seconds) * time.Second)
	}

}
