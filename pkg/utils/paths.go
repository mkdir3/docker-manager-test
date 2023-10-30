package utils

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func ResolveHomeDir(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}

func GetComposeFilePath(projectDir string) (string, error) {
	ymlPath := projectDir + "/docker-compose.yml"
	if _, err := os.Stat(ymlPath); err == nil {
		return ymlPath, nil
	}
	yamlPath := projectDir + "/docker-compose.yaml"
	if _, err := os.Stat(yamlPath); err == nil {
		return yamlPath, nil
	}
	return "", fmt.Errorf("no docker-compose file found in %s", projectDir)
}
