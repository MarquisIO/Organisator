package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	validator "gopkg.in/validator.v2"
)

// Configuration represents the configuration file of the project.
type Configuration struct {
	Labels       []Label `json:"labels"`
	Organisation string  `json:"organisation" validate:"nonzero"`
}

// FromFileAt creates configuration object from a file at path.
// Returns nil if it could not open the file.
func FromFileAt(path *string) (*Configuration, error) {

	var config Configuration
	var labelsNames = make(map[string]bool)

	err := getData(path, &config)
	if err != nil {
		return nil, err
	}

	errorInLabels := false

	// Validate current JSON
	if err := validator.Validate(config); err != nil {
		fmt.Printf("Error in configuration : %v\n", err)
		errorInLabels = true
	}

	// Validate the JSON labels
	for index, label := range config.Labels {
		if err := validator.Validate(label); err != nil {
			fmt.Printf("Error in label %v : %v\n", index+1, err)
			errorInLabels = true
		}
		if labelsNames[label.Name] {
			fmt.Printf("Error in label %v : Label name already present.\n", index+1)
			errorInLabels = true
		}

		labelsNames[label.Name] = true
	}

	if errorInLabels {
		return nil, fmt.Errorf("Labels configuration is wrong.")
	}

	return &config, nil
}

// Opens file, read it and Unmarshal JSON data
func getData(path *string, config *Configuration) error {

	if path == nil {
		return fmt.Errorf("Path invalid.")
	}

	file, err := os.Open(*path)
	if err != nil {
		return fmt.Errorf("Error opening file : %v", err)
	}

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file : %v", err)
		return fmt.Errorf("Error reading file : %v", err)
	}

	if err := json.Unmarshal(jsonData, &config); err != nil {
		log.Fatalf("Error JSON Unmarshal : %v", err)
		return fmt.Errorf("Error JSON Unmarshal : %v", err)
	}

	return nil
}
