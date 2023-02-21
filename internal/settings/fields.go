package settings

import (
	"encoding/json"
	"errors"
	"os"

	"qt-proxyserver/internal/logger"
	"qt-proxyserver/uigen"
)

const fileFields = "fields.json"

type Field struct {
	ProxyAddress   string
	ServerLogin    string
	ServerPassword string
	ServerPort     string
}

func Fields(w *uigen.UIWindow, tEditLogger *logger.TEditLogger) {
	loadFields(w, tEditLogger)
	saveFields(w, tEditLogger)
}

func loadFields(w *uigen.UIWindow, tEditLogger *logger.TEditLogger) {
	if _, err := os.Stat(fileFields); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return
		}

		tEditLogger.InsertText(logger.Fatal, "file with fields", err.Error())

		return
	}

	fieldsData, err := os.ReadFile(fileFields)
	if err != nil {
		tEditLogger.InsertText(logger.Error, "settings: read file", err.Error())
	}

	var fields Field
	if err := json.Unmarshal(fieldsData, &fields); err != nil {
		tEditLogger.InsertText(logger.Error, "settings: json unmarshal", err.Error())
	}

	w.LineEditProxyAddress.SetText(fields.ProxyAddress)
	w.LineEditServerLogin.SetText(fields.ServerLogin)
	w.LineEditServerPassword.SetText(fields.ServerPassword)
	w.LineEditServerPort.SetText(fields.ServerPort)
}

func saveFields(w *uigen.UIWindow, tEditLogger *logger.TEditLogger) {
	w.PushButtonSave.ConnectClicked(func(bool) {
		fields := Field{
			ProxyAddress:   w.LineEditProxyAddress.Text(),
			ServerLogin:    w.LineEditServerLogin.Text(),
			ServerPassword: w.LineEditServerPassword.Text(),
			ServerPort:     w.LineEditServerPort.Text(),
		}

		fieldsData, err := json.MarshalIndent(&fields, "", "  ")
		if err != nil {
			tEditLogger.InsertText(logger.Error, "settings: json marshal", err.Error())

			return
		}

		if err := os.WriteFile(fileFields, fieldsData, 0o600); err != nil {
			tEditLogger.InsertText(logger.Error, "settings: os: write file", err.Error())

			return
		}

		tEditLogger.InsertText(logger.Info, "fields save", "ok")
	})
}
