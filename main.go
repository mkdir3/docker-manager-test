package main

import (
	"docker-manager/cmd"
	"docker-manager/pkg/utils"
)

func main() {
	utils.ProjectInfo()
	cmd.Execute()
}
