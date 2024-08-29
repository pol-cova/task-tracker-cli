package fileManager

import (
	"encoding/json"
	"errors"
	"os"
)

func SaveTasks(file string, tasks interface{}) error {
	// Check if file exists and create it if it doesn't
	if _, err := os.Stat(file); os.IsNotExist(err) {
		_, err := os.Create(file)
		if err != nil {
			return errors.New("error creating file")
		}
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return errors.New("error marshalling data")
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return errors.New("error writing file")
	}

	return nil
}

func ReadTasks(file string, tasks interface{}) error {
	// Check if file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil
	}

	// Read file
	data, err := os.ReadFile(file)
	if err != nil {
		return errors.New("error reading file")
	}

	// Unmarshal data
	err = json.Unmarshal(data, tasks)
	if err != nil {
		return errors.New("error unmarshalling data")
	}

	return nil
}
