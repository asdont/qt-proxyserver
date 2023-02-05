package settings

import (
	"github.com/therecipe/qt/widgets"

	"qt-proxyserver/uigen"
)

func Settings(w *uigen.UIWindow) {
	pushButtonStartDialogWhy(w.PushButtonWhy)
}

func pushButtonStartDialogWhy(pushButton *widgets.QPushButton) {
	pushButton.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(
			nil, "Информация", "Не сделано.", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
}
