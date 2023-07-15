package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// DISTROS SCRIPTS

func DistroScript(Distro, ScriptPath string) {
	var answer string
	fmt.Printf("Are you sure you want to install %s script? [Y/n]: ", Distro)
	fmt.Scan(&answer)

	switch answer {
	case "Y", "y", "Yes", "yes", "yeah", "ДА", "д", "да", "Да":
		c := exec.Command("sh", ScriptPath)
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
		c.Run()
	default:
		Clear()
		Logo()
		DistroMenu()
	}
}

// Russian

func RuDistroScript(Distro, ScriptPath string) {
	var answer string
	fmt.Printf("Вы уверены, что хотите выполнить скрипт установки %s ? [Да/нет]: ", Distro)
	fmt.Scan(&answer)

	switch answer {
	case "Y", "y", "Yes", "yes", "yeah", "ДА", "д", "Д", "да", "Да":
		c := exec.Command("sh", ScriptPath)
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
		c.Run()
	default:
		Clear()
		Logo()
		DistroMenu()
	}
}
