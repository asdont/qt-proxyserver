// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type UIWindow struct {
	CentralWidget            *widgets.QWidget
	TabWidgetServer          *widgets.QTabWidget
	TabServer                *widgets.QWidget
	GroupBoxLog              *widgets.QGroupBox
	ScrollAreaLog            *widgets.QScrollArea
	ScrollAreaWidgetContents *widgets.QWidget
	GroupBoxServer           *widgets.QGroupBox
	LineEditProxyAddress     *widgets.QLineEdit
	PushButtonStartServer    *widgets.QPushButton
	Tab                      *widgets.QWidget
	GroupBox                 *widgets.QGroupBox
	PushButtonWhy            *widgets.QPushButton
	Statusbar                *widgets.QStatusBar
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
	this.GroupBoxLog.SetGeometry(core.NewQRect4(13, 90, 550, 320))
	this.ScrollAreaLog = widgets.NewQScrollArea(this.GroupBoxLog)
	this.ScrollAreaLog.SetObjectName("ScrollAreaLog")
	this.ScrollAreaLog.SetGeometry(core.NewQRect4(6, 30, 540, 280))
	this.ScrollAreaLog.SetCursor(gui.NewQCursor2(core.Qt__IBeamCursor))
	this.ScrollAreaLog.SetWidgetResizable(true)
	this.ScrollAreaWidgetContents = widgets.NewQWidget(this.ScrollAreaLog, core.Qt__Widget)
	this.ScrollAreaWidgetContents.SetObjectName("ScrollAreaWidgetContents")
	this.ScrollAreaWidgetContents.SetGeometry(core.NewQRect4(0, 0, 538, 278))
	this.ScrollAreaLog.SetWidget(this.ScrollAreaWidgetContents)
	this.GroupBoxServer = widgets.NewQGroupBox(this.TabServer)
	this.GroupBoxServer.SetObjectName("GroupBoxServer")
	this.GroupBoxServer.SetGeometry(core.NewQRect4(13, 10, 550, 70))
	this.LineEditProxyAddress = widgets.NewQLineEdit(this.GroupBoxServer)
	this.LineEditProxyAddress.SetObjectName("LineEditProxyAddress")
	this.LineEditProxyAddress.SetGeometry(core.NewQRect4(20, 30, 400, 28))
	this.LineEditProxyAddress.SetInputMethodHints(core.Qt__ImhNone)
	this.PushButtonStartServer = widgets.NewQPushButton(this.GroupBoxServer)
	this.PushButtonStartServer.SetObjectName("PushButtonStartServer")
	this.PushButtonStartServer.SetGeometry(core.NewQRect4(440, 30, 100, 28))
	this.TabWidgetServer.AddTab(this.TabServer, "")
	this.Tab = widgets.NewQWidget(this.TabWidgetServer, core.Qt__Widget)
	this.Tab.SetObjectName("Tab")
	this.GroupBox = widgets.NewQGroupBox(this.Tab)
	this.GroupBox.SetObjectName("GroupBox")
	this.GroupBox.SetGeometry(core.NewQRect4(10, 10, 551, 401))
	this.PushButtonWhy = widgets.NewQPushButton(this.GroupBox)
	this.PushButtonWhy.SetObjectName("PushButtonWhy")
	this.PushButtonWhy.SetGeometry(core.NewQRect4(20, 30, 86, 27))
	this.TabWidgetServer.AddTab(this.Tab, "")
	Window.SetCentralWidget(this.CentralWidget)
	this.Statusbar = widgets.NewQStatusBar(Window)
	this.Statusbar.SetObjectName("Statusbar")
	Window.SetStatusBar(this.Statusbar)

	this.RetranslateUi(Window)
	this.TabWidgetServer.SetCurrentIndex(0)
}

func (this *UIWindow) RetranslateUi(Window *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	Window.SetWindowTitle(_translate("Window", "Прокси SOCKS5", "", -1))
	this.GroupBoxLog.SetTitle(_translate("Window", "Информация", "", -1))
	this.GroupBoxServer.SetTitle(_translate("Window", "Запуск прокси SOCKS5", "", -1))
	this.LineEditProxyAddress.SetPlaceholderText(_translate("Window", "socks5://login:password@1.2.3.4:9090", "", -1))
	this.PushButtonStartServer.SetText(_translate("Window", "Запуск", "", -1))
	this.TabWidgetServer.SetTabText(this.TabWidgetServer.IndexOf(this.TabServer), _translate("Window", "Сервер", "", -1))
	this.GroupBox.SetTitle(_translate("Window", "Здесь нет настроек", "", -1))
	this.PushButtonWhy.SetText(_translate("Window", "Почему?", "", -1))
	this.TabWidgetServer.SetTabText(this.TabWidgetServer.IndexOf(this.Tab), _translate("Window", "Настройки", "", -1))
}
