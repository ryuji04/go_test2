package main

func main2() {
	a := app.New()
	w := a.NewWindow("app")
	a.Settings().SetTheme(theme.DarkTheme())
	edit := widget.NewMultiLineEntry()
	sc := widget.NewScrollContainer(edit)
	fnd := widget.Newentry()
	inf := widget.NewLabel("information bar")

	//show alert
	showInfo := func(s string) {
		inf.SetText(s)
		dialog.showInformation("info", s, w)
	}

}
