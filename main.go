package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	textDisplay := tview.NewTextView().SetDynamicColors(true).SetRegions(false).SetScrollable(true)
	textDisplay.SetTitle("f1").SetBorder(true).SetTitleAlign(tview.AlignLeft)

	textArea := tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	textArea.SetBorder(true).SetTitleAlign(tview.AlignLeft).SetTitle("f2")

	regexArea := tview.NewTextArea().
		SetPlaceholder("Expression...")
	regexArea.SetBorder(true).SetTitle("f3").SetTitleAlign(tview.AlignLeft)

	pages := tview.NewPages()

	matchesDisplay := tview.NewList().ShowSecondaryText(false)
	matchesDisplay.SetBorder(true).SetTitle("Matches").
		SetTitleAlign(tview.AlignLeft).SetTitle("f4")

	getMatch := func() {
		matchesDisplay.Clear()
		str := textArea.GetText()
		expr := regexArea.GetText()

		if expr == "" {
			textDisplay.SetText(str)
			return
		}

		regex, err := regexp.Compile(expr)

		if err != nil {
			textDisplay.SetText(str)
			return
		}

		var fmtMatchedString string

		matchList := regex.FindAllString(str, -1)

		if matchList != nil {
			fmtMatchedString = str
			for _, v := range matchList {
				fmtMatchedString = strings.Replace(fmtMatchedString, v, fmt.Sprintf("[green]%s[white]", v), 1)
				matchesDisplay.AddItem(
					v, "", 'M', nil)
			}
		} else {
			textDisplay.SetText(str)
			return
		}

		textDisplay.SetText(fmtMatchedString)

	}

	regexArea.SetChangedFunc(getMatch)
	textArea.SetChangedFunc(getMatch)

	mainView := tview.NewGrid().
		SetRows(0, 0, 0, 0, 0, 0, 0, 3).
		AddItem(textDisplay, 0, 0, 5, 5, 0, 0, false).
		AddItem(textArea, 5, 0, 2, 5, 0, 0, false).
		AddItem(regexArea, 7, 0, 1, 5, 0, 0, true).
		AddItem(matchesDisplay, 0, 5, 8, 2, 0, 0, false)

	pages.AddAndSwitchToPage("main", mainView, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF1:
			app.SetFocus(textDisplay)
		case tcell.KeyF2:
			app.SetFocus(textArea)
		case tcell.KeyF3:
			app.SetFocus(regexArea)
		case tcell.KeyF4:
			app.SetFocus(matchesDisplay)
		}
		return event
	})

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
