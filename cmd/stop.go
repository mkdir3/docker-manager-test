package cmd

import (
	"docker-manager/pkg/docker"
	"docker-manager/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [project]",
	Short: "Stop a Docker project",
	Long:  `Stop a Docker project by its name`,
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
		err = docker.ExecuteDockerComposeCommand(projectDir, "down")
		if err != nil {
			fmt.Printf("Failed to stop project %s: %v\n", projectName, err)
			return
		}
		fmt.Printf("Stopped project %s\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
