package utils

import (
	"fmt"
	"os"
	"os/exec"
)

/*
* Here is Soft storage
* MyOs, Neofetch
 */

// os info
func MyOs() {
	cmd := exec.Command("neofetch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start Neofetch. %s\nInstalling neofetch...\n", err.Error())
		Neofetch()
	}
}
func Neofetch() {
	cmd := exec.Command("sh", "src/neofetch.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
	Clear()
	MyOs()
}
