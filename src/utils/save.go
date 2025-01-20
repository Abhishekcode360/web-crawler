package utils

import (
	"encoding/json"
	"os"
)

func SaveToFile(filename string, data ProductURLs) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
