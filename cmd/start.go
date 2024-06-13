package cmd

import (
	"docker-manager/pkg/docker"
	"docker-manager/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [project]",
	Short: "Start a Docker project",
	Long:  `Start all Docker containers of a project`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath, ok := docker.Projects[projectName]
		if !ok {
			fmt.Printf("Unknown project: %s\n", projectName)
			return
		}
		projectDir, err := utils.ResolveHomeDir(projectPath)
		if err != nil {
			fmt.Printf("Failed to resolve home directory in %s: %v\n", projectPath, err)
			return
		}
		err = docker.ExecuteDockerComposeCommand(projectDir, "up", "-d --remove-orphans")
		if err != nil {
			fmt.Printf("Failed to start project %s: %v\n", projectName, err)
			return
		}
		fmt.Printf("Started project %s\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
