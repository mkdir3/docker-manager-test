package docker

import (
	"docker-manager/pkg/utils"
	"os"
	"os/exec"
)

func ExecuteDockerComposeCommand(projectDir string, args ...string) error {
	composeFilePath, err := utils.GetComposeFilePath(projectDir)
	if err != nil {
		return err
	}
	commandArgs := append([]string{"compose", "-f", composeFilePath}, args...)
	command := exec.Command("docker", commandArgs...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
