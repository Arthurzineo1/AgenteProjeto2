package logging

import (
	"log"
	"os"
)

var (
	infoLogger   *log.Logger
	errorLogger  *log.Logger
	infoLogFile  *os.File
	errorLogFile *os.File
)

// Init initializes the loggers
func Init() {
	var err error

	infoLogFile, err = os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Falha ao abrir o arquivo de log de informações:", err)
	}

	errorLogFile, err = os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Falha ao abrir o arquivo de log de erros:", err)
	}

	infoLogger = log.New(infoLogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorLogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an information message
func Info(message string) {
	infoLogger.Println(message)
}

// Error logs an error message
func Error(err error) {
	errorLogger.Println(err)
}

// Debug logs a debug message
func Debug(message string) {
	infoLogger.Println("DEBUG: ", message)
}

// Close closes the log files
func Close() {
	if infoLogFile != nil {
		infoLogFile.Close()
	}
	if errorLogFile != nil {
		errorLogFile.Close()
	}
}
