package data

import (
	"os"
	"os/user"
	"strings"
)

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return hostname, nil
}

func GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	username := currentUser.Username
	if strings.Contains(username, "\\") {
		parts := strings.Split(username, "\\")
		if len(parts) > 0 {
			return parts[len(parts)-1], nil
		}
	}

	return username, nil
}
