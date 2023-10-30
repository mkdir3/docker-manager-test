package docker

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
)

var Projects = make(map[string]string)

func init() {
	// Load projects from file when the package is initialized
	if err := LoadProjectsFromFile("projects.json"); err != nil {
		// Handle error, e.g., log it or print it
		// For now, let's just initialize an empty Projects map
		Projects = make(map[string]string)
	}
}

func GetSortedProjectNames() []string {
	projectNames := make([]string, 0, len(Projects))
	for projectName := range Projects {
		projectNames = append(projectNames, projectName)
	}
	sort.Strings(projectNames)
	return projectNames
}

func AddProject() {
	var projectName, projectPath string
	survey.AskOne(&survey.Input{Message: "Enter the project name:"}, &projectName)
	survey.AskOne(&survey.Input{Message: "Enter the project path:"}, &projectPath)

	var confirm string
	confirmPrompt := &survey.Select{
		Message: fmt.Sprintf("Do you want to add project %s with path %s?", projectName, projectPath),
		Options: []string{"Yes", "No"},
	}
	survey.AskOne(confirmPrompt, &confirm)

	if confirm == "Yes" {
		Projects[projectName] = projectPath
		SaveProjectsToFile("projects.json")	
		fmt.Printf("Added project %s\n", projectName)
	}
}

func RemoveProject() error {
	var projectName string
	survey.AskOne(&survey.Input{Message: "Enter the project name you'd like to remove:"}, &projectName)

	if _, exists := Projects[projectName]; !exists {
		return fmt.Errorf("Project %s does not exist", projectName)
	}

	delete(Projects, projectName)

	if err := SaveProjectsToFile("projects.json"); err != nil {
		return fmt.Errorf("Failed to save projects after removal: %v", err)
	}

	return nil
}