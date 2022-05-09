package main

import (
	"fmt"
	"os/exec"

	"github.com/mbndr/figlet4go"
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

func versionf() {
	fmt.Println("GOAPPLE 2.0")
}

func descr() {
	fmt.Print(" \t\t\t\033[33mversion \033[31m2.1\033[0m\n\n USAGE:\n\n\033[1m goapple-cli\033[0m\t to start CLI application\n\n\033[1m goapple-gui\033[0m\t to start GUI application\n\n")
}

func main() {
	clear()
	logo()
	descr()
}
