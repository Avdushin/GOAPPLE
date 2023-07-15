package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/logrusorgru/aurora" // Colors
	"github.com/mbndr/figlet4go"      // Logotype
)

var (
	version = "2.3"
	scaner  string
)

func clear() {
	out, _ := exec.Command("clear").Output()
	fmt.Println(string(out))
}

func logo() {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}
	renderStr, _ := ascii.RenderOpts("GOAPPLE", options)
	fmt.Print(renderStr)
}

// Exec commands
func ExecScript(command string) error {
	args := strings.Split(command, " ")
	if len(args) < 1 {
		return fmt.Errorf("command is empty")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}

func descr() {
	fmt.Println(fmt.Sprint(Bold(Yellow("\t\t\tversion")), Red(version), Bold("\n\n USAGE:\n\n"), Bold("goapple-cli"), "\t\t to start CLI application\n\n", Bold(" goapple-gui"), "\t\t to start GUI application\n\n"))
}

func input() {
	fmt.Print(" Type \"cli\", \"gui\" or \"quit\" > ")

	fmt.Scan(&scaner)
	switch scaner {
	case "goapple-cli", "cli", "1":
		if err := ExecScript("goapple-cli"); err != nil {
			fmt.Println("Ошибка выполнения команды 'goapple-cli':", err)
		}
	case "goapple-gui", "gui", "2":
		if err := ExecScript("goapple-gui"); err != nil {
			fmt.Println("Ошибка выполнения команды 'goapple-gui':", err)
		}
	case "quit", "q", "exit", "0":
		return
	default:
		clear()
		logo()
		descr()
		input()
	}
}
func main() {
	clear()
	logo()
	descr()
	input()
}
