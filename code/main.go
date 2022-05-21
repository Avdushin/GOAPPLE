package main

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/logrusorgru/aurora" // Colors
	"github.com/mbndr/figlet4go"      // Logotype
)

var (
	version = "2.2"
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

func descr() {
	fmt.Println(fmt.Sprint(Bold(Yellow("\t\t\tversion")), Red(version), Bold("\n\n USAGE:\n\n"), Bold("goapple-cli"), "\t\t to start CLI application\n\n", Bold(" goapple-gui"), "\t\t to start GUI application\n\n"))
}

func cli() {
	cmd := exec.Command("goapple-cli")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func gui() {
	cmd := exec.Command("goapple-gui")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func input() {
	fmt.Print(" Type \"cli\", \"gui\" or \"quit\" => ")

	fmt.Scan(&scaner)
	if scaner == "goapple-cli" {
		cli()
	} else if scaner == "cli" {
		cli()
	} else if scaner == "goapple-gui" {
		gui()
	} else if scaner == "gui" {
		gui()
	} else if scaner == "quit" {
		return
	} else if scaner == "exit" {
		return
	} else if scaner == "0" {
		return
	} else {
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
