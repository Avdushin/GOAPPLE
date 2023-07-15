package main

import (
	"flag"

	"github.com/avdushin/pkg/utils"
)

func main() {
	lang := flag.Bool("ru", false, "Use to start russian version")
	flag.Parse()
	LangCheck := *lang
	// -ru flag check
	switch LangCheck {
	case true:
		utils.Clear()
		utils.Logo()
		utils.RuApp()
	default:
		utils.Clear()
		utils.Logo()
		utils.MainMenu()
	}
}
