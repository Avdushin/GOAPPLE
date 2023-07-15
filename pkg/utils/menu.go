package utils

import (
	"fmt"
	"os"

	"github.com/avdushin/pkg/consts"
)

// Menus
// main menu
func MainMenu() {
	fmt.Printf(" \033[36m1)\033[37m DISTROS  \033[36m2)\033[37m MYOS\n\n \033[36m3)\033[37m SETTINGS \033[36m\033[36m0)\033[31m EXIT\n\n\033[0m")

	var scan string

	fmt.Print("\033[33m SELECT: \033[0m")
	fmt.Scan(&scan)

	switch scan {
	case "0", "exit", "quit":
		os.Exit(0)
	case "1":
		Clear()
		Logo()
		DistroMenu()
	case "2":
		Clear()
		MyOs()
		MainMenu()
	case "3":
		Clear()
		Logo()
		Settings()
	case "ru", "rus":
		Clear()
		Logo()
		RuApp()
	default:
		Clear()
		Logo()
		MainMenu()
	}
}

// Distros menu
func DistroMenu() {
	fmt.Printf(" \033[36m1)\033[37m BSPWM+I3WM \033[36m\n 2)\033[37m BSPWM\n\033[36m 3)\033[37m I3-WM\n\n\033[36m\033[36m\033[36m 5)\033[33m BACK \033[36m0)\033[31m EXIT\n\n\033[0m")

	var scan string

	fmt.Printf("\033[33m %s: \033[0m", consts.Selecten)
	fmt.Scan(&scan)

	switch scan {
	case "0", "exit", "EXIT", "q", "quit", "QUIT", "выход", "ВЫХОД":
		os.Exit(0)
	case "1":
		DistroScript(
			"BSPWM+I3",
			"src/distros/setup.sh",
		)
	case "2":
		DistroScript(
			"I3-WM",
			"src/distros/i3/setup.sh",
		)
	case "3":
		DistroScript(
			"BSPWM",
			"src/distros/bspwm/setup.sh",
		)
	case "5":
		Clear()
		Logo()
		MainMenu()
	case "test":
		DistroScript("Test", "src/scripts/test.sh")
	default:
		Clear()
		Logo()
		DistroMenu()
	}
}

// settings
func Settings() {
	fmt.Printf("\n \033[33mAUTHOR:\033[0m https://github.com/Avdushin\n \033[33mVERSION: \033[0m%sgo\033[36m \n\n \033[33mLANGUAGE:\n\n \033[36m1)\033[0m РУССКИЙ \n \033[36m2)\033[0m ENGLISH \n\n \033[36m5) \033[33mBACK \033[36m0)\033[31m EXIT\n\n", consts.Version)

	var scan string

	fmt.Print("\n\033[33m ACTION: \033[0m")
	fmt.Scan(&scan)

	switch scan {
	case "0", "exit", "EXIT", "q", "quit", "QUIT", "выход", "ВЫХОД":
		os.Exit(0)
	case "1":
		Clear()
		Logo()
		RuApp()
	case "2":
		Clear()
		Logo()
		MainMenu()
	case "5":
		Clear()
		Logo()
		MainMenu()
	default:
		Clear()
		Logo()
		Settings()
	}
}
