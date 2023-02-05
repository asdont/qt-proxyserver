package server

import (
	"github.com/therecipe/qt/widgets"

	"qt-proxyserver/uigen"
)

func Server(w *uigen.UIWindow) {
	StartServer(w.PushButtonStartServer)
}

func StartServer(pushButton *widgets.QPushButton) {
	pushButton.ConnectClicked(func(bool) {

	})
}
