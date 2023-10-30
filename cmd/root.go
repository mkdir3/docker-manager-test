package cmd

import (
	"docker-manager/pkg/docker"
	"docker-manager/pkg/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-manager",
	Short: "CLI app to manage Dockerized projects",
	Long:  `A CLI app to manage Dockerized projects. It can start, stop and list running Docker containers.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := docker.CheckAndLoadProjectsFile("projects.json"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		selectedProjects, err := docker.SelectProjects()
		if err != nil {
			fmt.Printf("Failed to select projects: %v\n", err)
			return
		}

		for _, projectName := range selectedProjects {
			projectPath, ok := docker.Projects[projectName]
			if !ok {
				fmt.Printf("Unknown project: %s\n", projectName)
				continue
			}
			projectDir, err := utils.ResolveHomeDir(projectPath)
			if err != nil {
				fmt.Printf("Failed to resolve home directory in %s: %v\n", projectPath, err)
				continue
			}
			err = docker.ExecuteDockerComposeCommand(projectDir, "up", "-d")
			if err != nil {
				fmt.Printf("Failed to start project %s: %v\n", projectName, err)
				continue
			}
			fmt.Printf("Started project %s\n", projectName)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
