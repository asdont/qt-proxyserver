package uicustom

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"qt-proxyserver/uigen"
)

func SetupWindowCustom(widget *widgets.QMainWindow) {
	widget.SetGeometry(core.NewQRect4(200, 150, 585, 475))
	widget.SetFixedSize2(585, 475)
}

func SetupUICustom(ui *uigen.UIWindow) {}
