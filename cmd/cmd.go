package cmd

import (
	"os"
	"os/exec"
)

func ClearTerminal() (err error) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return
	}
	return
}
