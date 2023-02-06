package logger

import (
	"fmt"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Level string

const (
	Fatal Level = "FATAL"
	Error       = "ERROR"
	Info        = "INFO"
)

const bufChanTEditLogger = 100

type TEditLogger struct {
	msg chan string
}

func NewTEditLogger(w *widgets.QTextEdit) *TEditLogger {
	tEditLogger := new(TEditLogger)

	tEditLogger.msg = make(chan string, bufChanTEditLogger)

	go func() {
		for {
			w.InsertPlainText(<-tEditLogger.msg)
			w.MoveCursor(gui.QTextCursor__End, gui.QTextCursor__MoveAnchor)
		}
	}()

	return tEditLogger
}

func (tEL *TEditLogger) Write(p []byte) (n int, err error) {
	tEL.msg <- string(p)

	return len(p), nil
}

func (tEL *TEditLogger) InsertText(lvl Level, msg, text string) {
	tEL.msg <- fmt.Sprintf("[%s] %s: %s\n", lvl, msg, text)
}
