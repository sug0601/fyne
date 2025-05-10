package main

import (
	"fyne-app/launcher"
	"runtime"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("プロジェクト起動")

	button := widget.NewButton("プロジェクトを起動", func() {
		go launcher.LaunchProjects()
	})
	button.Importance = widget.HighImportance

	if runtime.GOOS != "darwin" {
		button.Disable()
		button.SetText("macOS 専用機能")
	} else if len(launcher.Projects) == 0 {
		button.Disable()
		button.SetText("プロジェクト未設定 (launcher/launcher.go を更新)")
	}
	w.SetContent(container.NewVBox(button))

	w.ShowAndRun()
}
