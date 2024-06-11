package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Создание приложения
	myApp := app.New()

	// Создание окна
	myWindow := myApp.NewWindow("Пример приложения")

	// Добавление виджетов
	hello := widget.NewLabel("Привет, мир!")
	myWindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("OK", func() {
			hello.SetText("Добро пожаловать в Fyne!")
		}),
	))

	// Задать размер окна
	myWindow.Resize(fyne.NewSize(200, 150))

	// Отображение окна
	myWindow.ShowAndRun()
}
