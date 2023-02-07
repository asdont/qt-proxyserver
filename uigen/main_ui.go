// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIWindow struct {
	CentralWidget          *widgets.QWidget
	TabWidgetServer        *widgets.QTabWidget
	TabServer              *widgets.QWidget
	GroupBoxLog            *widgets.QGroupBox
	TextEditLog            *widgets.QTextEdit
	GroupBoxServer         *widgets.QGroupBox
	LineEditProxyAddress   *widgets.QLineEdit
	PushButtonServerStart  *widgets.QPushButton
	LineEditServerLogin    *widgets.QLineEdit
	LineEditServerPassword *widgets.QLineEdit
	LineEditServerPort     *widgets.QLineEdit
	PushButtonSave         *widgets.QPushButton
	Statusbar              *widgets.QStatusBar
}

func (this *UIWindow) SetupUI(Window *widgets.QMainWindow) {
	Window.SetObjectName("Window")
	Window.SetEnabled(true)
	Window.SetGeometry(core.NewQRect4(0, 0, 585, 475))
	this.CentralWidget = widgets.NewQWidget(Window, core.Qt__Widget)
	this.CentralWidget.SetObjectName("CentralWidget")
	this.TabWidgetServer = widgets.NewQTabWidget(this.CentralWidget)
	this.TabWidgetServer.SetObjectName("TabWidgetServer")
	this.TabWidgetServer.SetGeometry(core.NewQRect4(1, 0, 580, 450))
	this.TabServer = widgets.NewQWidget(this.TabWidgetServer, core.Qt__Widget)
	this.TabServer.SetObjectName("TabServer")
	this.GroupBoxLog = widgets.NewQGroupBox(this.TabServer)
	this.GroupBoxLog.SetObjectName("GroupBoxLog")
	this.GroupBoxLog.SetGeometry(core.NewQRect4(13, 119, 550, 291))
	this.TextEditLog = widgets.NewQTextEdit(this.GroupBoxLog)
	this.TextEditLog.SetObjectName("TextEditLog")
	this.TextEditLog.SetGeometry(core.NewQRect4(10, 30, 530, 250))
	this.TextEditLog.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOn)
	this.GroupBoxServer = widgets.NewQGroupBox(this.TabServer)
	this.GroupBoxServer.SetObjectName("GroupBoxServer")
	this.GroupBoxServer.SetGeometry(core.NewQRect4(13, 10, 550, 101))
	this.LineEditProxyAddress = widgets.NewQLineEdit(this.GroupBoxServer)
	this.LineEditProxyAddress.SetObjectName("LineEditProxyAddress")
	this.LineEditProxyAddress.SetGeometry(core.NewQRect4(20, 26, 400, 28))
	this.LineEditProxyAddress.SetInputMethodHints(core.Qt__ImhNone)
	this.PushButtonServerStart = widgets.NewQPushButton(this.GroupBoxServer)
	this.PushButtonServerStart.SetObjectName("PushButtonServerStart")
	this.PushButtonServerStart.SetEnabled(true)
	this.PushButtonServerStart.SetGeometry(core.NewQRect4(440, 26, 100, 28))
	this.LineEditServerLogin = widgets.NewQLineEdit(this.GroupBoxServer)
	this.LineEditServerLogin.SetObjectName("LineEditServerLogin")
	this.LineEditServerLogin.SetGeometry(core.NewQRect4(20, 63, 170, 28))
	this.LineEditServerPassword = widgets.NewQLineEdit(this.GroupBoxServer)
	this.LineEditServerPassword.SetObjectName("LineEditServerPassword")
	this.LineEditServerPassword.SetGeometry(core.NewQRect4(195, 63, 170, 28))
	this.LineEditServerPassword.SetEchoMode(widgets.QLineEdit__Password)
	this.LineEditServerPort = widgets.NewQLineEdit(this.GroupBoxServer)
	this.LineEditServerPort.SetObjectName("LineEditServerPort")
	this.LineEditServerPort.SetGeometry(core.NewQRect4(370, 63, 50, 28))
	this.LineEditServerPort.SetMaxLength(5)
	this.PushButtonSave = widgets.NewQPushButton(this.GroupBoxServer)
	this.PushButtonSave.SetObjectName("PushButtonSave")
	this.PushButtonSave.SetEnabled(true)
	this.PushButtonSave.SetGeometry(core.NewQRect4(440, 63, 100, 28))
	this.TabWidgetServer.AddTab(this.TabServer, "")
	Window.SetCentralWidget(this.CentralWidget)
	this.Statusbar = widgets.NewQStatusBar(Window)
	this.Statusbar.SetObjectName("Statusbar")
	Window.SetStatusBar(this.Statusbar)

	this.RetranslateUi(Window)
	this.TabWidgetServer.SetCurrentIndex(0)
}

func (this *UIWindow) RetranslateUi(Window *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	Window.SetWindowTitle(_translate("Window", "Сервер через SOCKS5", "", -1))
	this.GroupBoxLog.SetTitle(_translate("Window", "Информация", "", -1))
	this.TextEditLog.SetWhatsThis(_translate("Window", "Инормация. Сюда пишутся ошибки и логи.", "", -1))
	this.GroupBoxServer.SetTitle(_translate("Window", "Запуск сервера через прокси SOCKS5", "", -1))
	this.LineEditProxyAddress.SetWhatsThis(_translate("Window", "Серверные запросы через этот прокси", "", -1))
	this.LineEditProxyAddress.SetPlaceholderText(_translate("Window", "socks5://login:password@1.2.3.4:9090", "", -1))
	this.PushButtonServerStart.SetText(_translate("Window", "Запустить", "", -1))
	this.LineEditServerLogin.SetWhatsThis(_translate("Window", "Логин для сервера", "", -1))
	this.LineEditServerLogin.SetText(_translate("Window", "", "", -1))
	this.LineEditServerLogin.SetPlaceholderText(_translate("Window", "Логин", "", -1))
	this.LineEditServerPassword.SetWhatsThis(_translate("Window", "Пароль для сервера", "", -1))
	this.LineEditServerPassword.SetText(_translate("Window", "", "", -1))
	this.LineEditServerPassword.SetPlaceholderText(_translate("Window", "Пароль", "", -1))
	this.LineEditServerPort.SetWhatsThis(_translate("Window", "Порт для сервера", "", -1))
	this.LineEditServerPort.SetPlaceholderText(_translate("Window", "Порт", "", -1))
	this.PushButtonSave.SetText(_translate("Window", "Сохранить", "", -1))
	this.TabWidgetServer.SetTabText(this.TabWidgetServer.IndexOf(this.TabServer), _translate("Window", "Сервер", "", -1))
}
