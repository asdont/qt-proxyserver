package uicustom

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"qt-proxyserver/uigen"
)

func FieldsMaskValidator(ui *uigen.UIWindow) {
	proxyAddress := core.NewQRegExp()
	serverLogin := core.NewQRegExp()
	serverPassword := core.NewQRegExp()
	serverPort := core.NewQRegExp()

	proxyAddress.SetPattern(`[^\s]+`)
	serverLogin.SetPattern(`[^\s]{1,22}`)
	serverPassword.SetPattern(`[^\s]{1,22}`)
	serverPort.SetPattern(`^\d{1,5}$`)

	ui.LineEditProxyAddress.SetValidator(gui.NewQRegExpValidator2(proxyAddress, nil))
	ui.LineEditServerLogin.SetValidator(gui.NewQRegExpValidator2(serverLogin, nil))
	ui.LineEditServerPassword.SetValidator(gui.NewQRegExpValidator2(serverPassword, nil))
	ui.LineEditServerPort.SetValidator(gui.NewQRegExpValidator2(serverPort, nil))
}
