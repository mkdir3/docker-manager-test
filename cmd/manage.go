package cmd

import (
	"docker-manager/pkg/docker"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage Docker projects",
	Long:  `Add or remove Docker projects paths.`,
	Run: func(cmd *cobra.Command, args []string) {
		var action string
		actionPrompt := &survey.Select{
			Message: "What do you want to do?",
			Options: []string{"Add", "Remove"},
		}
		survey.AskOne(actionPrompt, &action)

		switch action {
		case "Add":
			docker.AddProject()	

		case "Remove":
			docker.RemoveProject()
		}
	},
}

func init() {
	rootCmd.AddCommand(manageCmd)
}
