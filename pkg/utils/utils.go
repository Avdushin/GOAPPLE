package utils

import (
	"fmt"
	"os/exec"

	"github.com/mbndr/figlet4go"
)

// Clear screen
func Clear() {
	out, _ := exec.Command("clear").Output()
	fmt.Println(string(out))
}

// Logotype output
func Logo() {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	// options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}
	renderStr, _ := ascii.RenderOpts("PIN3APPLE", options)
	fmt.Print(renderStr)
}
