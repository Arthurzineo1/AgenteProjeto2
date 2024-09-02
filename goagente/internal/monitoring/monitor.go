package monitoring

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"goagente/internal/logging"
	"os"
)

// CreateHashFiles checks if the hash files exist, and if not, creates them.
func CreateHashFiles() {
	files := []string{"hashcore.txt", "hashprogram.txt", "hashhardware.txt"}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			// File does not exist, so create it
			f, err := os.Create(file)
			if err != nil {
				newErr := fmt.Errorf("error creating hash file %s: %v", file, err)
				logging.Error(newErr)
			}
			defer f.Close()
		}
	}
}

func CompareAndUpdateHashHardware(jsonData string) bool {
	// Compute the hash of the JSON string
	hash := sha256.Sum256([]byte(jsonData))
	hashString := hex.EncodeToString(hash[:])

	// Read the current hash from the hashhardware.txt file
	existingHash, err := os.ReadFile("hashhardware.txt")
	if err != nil && !os.IsNotExist(err) {
		logging.Error(err)
	}

	// Compare the existing hash with the new hash
	if string(existingHash) == hashString {
		// Hashes are equal, no update needed
		return false
	}

	// Hashes are different, update the file
	err = os.WriteFile("hashhardware.txt", []byte(hashString), 0644)
	if err != nil {
		logging.Error(err)
	}

	return true
}

func CompareAndUpdateHashCore(jsonData string) bool {
	// Compute the hash of the JSON string
	hash := sha256.Sum256([]byte(jsonData))
	hashString := hex.EncodeToString(hash[:])

	// Read the current hash from the hashhardware.txt file
	existingHash, err := os.ReadFile("hashcore.txt")
	if err != nil && !os.IsNotExist(err) {
		logging.Error(err)
	}

	// Compare the existing hash with the new hash
	if string(existingHash) == hashString {
		// Hashes are equal, no update needed
		return false
	}

	// Hashes are different, update the file
	err = os.WriteFile("hashcore.txt", []byte(hashString), 0644)
	if err != nil {
		logging.Error(err)
	}

	return true
}

func CompareAndUpdateHashPrograms(jsonData string) bool {
	// Compute the hash of the JSON string
	hash := sha256.Sum256([]byte(jsonData))
	hashString := hex.EncodeToString(hash[:])

	// Read the current hash from the hashhardware.txt file
	existingHash, err := os.ReadFile("hashprogram.txt")
	if err != nil && !os.IsNotExist(err) {
		logging.Error(err)
	}

	// Compare the existing hash with the new hash
	if string(existingHash) == hashString {
		// Hashes are equal, no update needed
		return false
	}

	// Hashes are different, update the file
	err = os.WriteFile("hashprogram.txt", []byte(hashString), 0644)
	if err != nil {
		logging.Error(err)
	}

	return true
}
