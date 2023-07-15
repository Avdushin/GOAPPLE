package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GOAPPLE-GUI")
	myWindow.Resize(fyne.NewSize(430, 370))

	// set header icon
	hicon, _ := fyne.LoadResourceFromPath("/usr/local/share/applications/src/assets/icons/logo.jpg")
	myApp.SetIcon(hicon)

	// Language settings
	// english
	fileLabel := "File"
	aboutLabel := "About"
	settingsLabel := "Settings"
	themeLabel := "Theme"
	darkThemeLabel := "Dark Theme"
	LightThemeLabel := "Light Theme"
	setupBTNLabel := "SETUP"
	errorDistroLabel := " ERROR: Select your distro!"
	errorLanguageLabel := " PLEASE RESTART THE APPLICATION!"
	langLabel := "Language"
	ruLangLabel := "Russian"
	enLangLabel := "English"
	label2Label := "Please select distro:"

	// header menu
	about := fyne.NewMenuItem(aboutLabel, func() {
		url := "https://github.com/avdushin"
		exec.Command("xdg-open", url).Start()
	})

	file_data, err := ioutil.ReadFile("/etc/goapple-gui/language.txt")
	if err != nil {
		log.Fatal("I can't to read :(\n Sry Man... One Second!")
		e := ioutil.WriteFile("/etc/goapple-gui/language.txt", []byte("true"), 0777)
		if e != nil {
			log.Fatal(e)
		}
	}

	var Languager bool // bool language

	switch string(file_data) {
	case "false":
		Languager = false
		log.Printf("English Language is %t", Languager)
		fileLabel = "Файл"
		aboutLabel = "О проекте"
		settingsLabel = "Настройки"
		themeLabel = "Тема"
		darkThemeLabel = "Темная тема"
		LightThemeLabel = "Светлая тема"
		setupBTNLabel = "УСТАНОВИТЬ"
		errorDistroLabel = " ОШИБКА: Вы не выбрали дистрибутив!"
		label2Label = "Выберите дистрибутив для установки:"
		errorLanguageLabel = " ПЕРЕЗАПУСТИТЕ ПРИЛОЖЕНИЕ,ЧТОБЫ СМЕНИТЬ ЯЗЫК!"
	case "true":
		Languager = true
		log.Printf("English Language is %t", Languager)
		fileLabel = "File"
		aboutLabel = "About"
		settingsLabel = "Settings"
		themeLabel = "Theme"
		darkThemeLabel = "Dark Theme"
		LightThemeLabel = "Light Theme"
		setupBTNLabel = "SETUP"
		errorDistroLabel = " ERROR Select your distro!"
		langLabel = "Language"
		ruLangLabel = "Russian"
		enLangLabel = "English"
		label2Label = "Please select distro:"
		errorLanguageLabel = " PLEASE RESTART THE APPLICATION!"
	}

	// set error's text color
	errorColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	errorLanguage := canvas.NewText(errorLanguageLabel, errorColor)
	errorLanguage.Hide()
	// lang Menu item
	lang := fyne.NewMenuItem(langLabel, nil)

	ruLang := fyne.NewMenuItem(ruLangLabel, func() {
		e := ioutil.WriteFile("/etc/goapple-gui/language.txt", []byte("false"), 0777)
		if err != nil {
			log.Fatal(e)
		}
		file_data, err := ioutil.ReadFile("/etc/goapple-gui/language.txt")
		if err != nil {
			log.Fatal("Can't to read language settings file")
		}
		log.Print(string(file_data))
		log.Print(string("/etc/goapple-gui/language.txt"))
		errorLanguage.Show()
	})
	enLang := fyne.NewMenuItem(enLangLabel, func() {
		e := ioutil.WriteFile("/etc/goapple-gui/language.txt", []byte("true"), 0777)
		if err != nil {
			log.Fatal(e)
		}
		file_data, err := ioutil.ReadFile("/etc/goapple-gui/language.txt")
		if err != nil {
			log.Fatal("Can't to read language settings file")
		}
		log.Print(string(file_data))
		errorLanguage.Show()
	})

	lang.ChildMenu = fyne.NewMenu(
		"",
		ruLang,
		enLang,
	)

	// Themes menu
	Theme := fyne.NewMenuItem(themeLabel, nil)

	dark := fyne.NewMenuItem(darkThemeLabel, func() { myApp.Settings().SetTheme(theme.DarkTheme()) })
	light := fyne.NewMenuItem(LightThemeLabel, func() { myApp.Settings().SetTheme(theme.LightTheme()) })

	Theme.ChildMenu = fyne.NewMenu(
		"",
		dark,
		light,
	)

	FileMenu := fyne.NewMenu(fileLabel, about)
	SettingsMenu := fyne.NewMenu(settingsLabel, Theme, lang)

	// main menu
	menu := fyne.NewMainMenu(FileMenu, SettingsMenu)
	// setup menu
	myWindow.SetMainMenu(menu)

	// progress bar
	infinite := widget.NewProgressBarInfinite()
	infinite.Stop()

	errorLabel2 := canvas.NewText(errorDistroLabel, errorColor)
	errorLabel2.Hide()

	selectDistro := widget.NewSelect(
		[]string{"Manjaro", "Solus", "Fedora"}, func(distro string) {
			log.Printf("Selected distro : %s", distro)
		})

	setupBTN := widget.NewButton(setupBTNLabel, func() {
		switch selectDistro.Selected {
		case "Manjaro":
			infinite.Start()
			c := exec.Command("st", "-f", "30", "sh", "/usr/local/share/applications/src/distros/Manjaro/manjaro.sh")
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Run()
		case "Solus":
			infinite.Start()
			c := exec.Command("st", "-f", "30", "sh", "/usr/local/share/applications/src/distros/Solus/solus.sh")
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Run()
		case "Fedora":
			infinite.Start()
			c := exec.Command("st", "-f", "30", "sh", "/usr/local/share/applications/src/distros/Fedora/fedora.sh")
			c.Stdout = os.Stdout
			c.Stdin = os.Stdin
			c.Stderr = os.Stderr
			c.Run()
		}
		// IsDistro Selected check
		switch selectDistro.Selected {
		case "":
			infinite.Stop()
			errorLabel2.Show()
		default:
			errorLabel2.Hide()
		}
	})

	// app label
	label2 := widget.NewLabel(label2Label)

	myWindow.SetContent(container.NewVBox(
		errorLanguage,
		label2,
		selectDistro,
		errorLabel2,
		setupBTN,
		infinite,
	))
	myWindow.ShowAndRun()
}
