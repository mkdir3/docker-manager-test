package docker

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func LoadProjectsFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal the JSON into the Projects map
	err = json.Unmarshal(data, &Projects)
	if err != nil {
		return err
	}

	return nil
}

// SaveProjectsToFile saves the Projects map to a JSON file
func SaveProjectsToFile(filename string) error {
	data, err := json.Marshal(Projects)
	if err != nil {
		return err
	}

	// Write the JSON to the file
	err = os.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func CheckAndLoadProjectsFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		var createFile string
		prompt := &survey.Select{
			Message: fmt.Sprintf("The project configuration file '%s' does not exist. Would you like to create one and add projects? 😎", filePath),
			Options: []string{"Yes", "No"},
		}
		survey.AskOne(prompt, &createFile)
		if createFile == "Yes" {
			AddProject()
			SaveProjectsToFile(filePath)
		} else {
			fmt.Println("No projects file found. Exiting.")
			os.Exit(1)
		}
	} else {
		// Load projects from file
		if err := LoadProjectsFromFile(filePath); err != nil {
			return fmt.Errorf("Failed to load projects: %v", err)
		}
	}
	return nil
}