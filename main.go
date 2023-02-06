package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"qt-proxyserver/internal/logger"
	"qt-proxyserver/internal/server"
	"qt-proxyserver/internal/settings"
	"qt-proxyserver/uicustom"
	"qt-proxyserver/uigen"
)

type Window struct {
	Widget *widgets.QMainWindow
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	w := NewServerWidget(nil)

	w.Widget.Show()

	os.Exit(app.Exec())
}

func NewServerWidget(parent widgets.QWidget_ITF) *Window {
	window := &Window{
		Widget: widgets.NewQMainWindow(parent, core.Qt__Window),
	}

	uiMainWindow := new(uigen.UIWindow)
	uiMainWindow.SetupUI(window.Widget)

	uicustom.SetupWindowCustom(window.Widget)
	uicustom.SetupUICustom(uiMainWindow)

	uicustom.FieldsMaskValidator(uiMainWindow)

	server.Server(uiMainWindow, server.CompileRegexps(), logger.NewTEditLogger(uiMainWindow.TextEditLog))
	settings.Settings(uiMainWindow)

	return window
}
