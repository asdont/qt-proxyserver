package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/therecipe/qt/widgets"
	"golang.org/x/net/proxy"

	"qt-proxyserver/internal/logger"
	"qt-proxyserver/internal/socks5"
	"qt-proxyserver/uigen"
)

const regexpSOCKS5 = `^socks5://(.+:.+@|)(\d{1,3}[.]){3}\d{1,3}:\d{1,5}$`

var (
	errStopServer        = errors.New("stop server")
	errWrongProxyAddress = errors.New("wrong proxy address")
)

func Server(w *uigen.UIWindow) {
	tEditLogger := logger.NewTEditLogger(w.TextEditLog)

	done := make(chan struct{}, 1)
	chErr := make(chan error, 1)

	go func() {
		for {
			err := <-chErr

			tEditLogger.InsertText(logger.Fatal, "global error", err.Error())
		}
	}()

	startServer(
		w.PushButtonServerStart,
		w.PushButtonServerStop,
		w.LineEditProxyAddress,
		w.LineEditserverLogin,
		w.LineEditServerPassword,
		w.LineEditServerPort,
		done,
		chErr,
		tEditLogger,
	)

	stopServer(
		w.PushButtonServerStart,
		w.PushButtonServerStop,
		strings.TrimSpace(w.LineEditServerPort.Text()),
		done,
		tEditLogger,
	)
}

func startServer(
	pushButtonStartServer *widgets.QPushButton,
	pushButtonStopServer *widgets.QPushButton,
	lineEditProxyAddress *widgets.QLineEdit,
	lineEditServerLogin *widgets.QLineEdit,
	lineEditServerPassword *widgets.QLineEdit,
	lineEditServerPort *widgets.QLineEdit,
	done chan struct{},
	chErr chan error,
	tEditLogger *logger.TEditLogger,
) {
	pushButtonStartServer.ConnectClicked(func(bool) {
		proxyAddress := strings.TrimSpace(lineEditProxyAddress.Text())

		if err := checkProxyAddress(proxyAddress); err != nil {
			tEditLogger.InsertText(logger.Error, "check proxy address", err.Error())

			return
		}

		pushButtonStartServer.SetEnabled(false)
		pushButtonStopServer.SetEnabled(true)

		tEditLogger.InsertText(logger.Info, "start server", lineEditServerPort.Text())

		if err := startSocks5(
			proxyAddress,
			strings.TrimSpace(lineEditServerLogin.Text()),
			strings.TrimSpace(lineEditServerPassword.Text()),
			strings.TrimSpace(lineEditServerPort.Text()),
			done,
			chErr,
			tEditLogger,
		); err != nil {
			tEditLogger.InsertText(logger.Fatal, "start socks5", err.Error())
		}
	})
}

func checkProxyAddress(proxyAddress string) error {
	if !regexp.MustCompile(regexpSOCKS5).MatchString(proxyAddress) {
		return fmt.Errorf("%s: %v", proxyAddress, errWrongProxyAddress)
	}

	return nil
}

func startSocks5(
	proxyAddress string,
	serverLogin string,
	serverPassword string,
	serverPort string,
	done chan struct{},
	chErr chan error,
	tEditLogger *logger.TEditLogger,
) error {
	credentials := socks5.StaticCredentials{
		serverLogin: serverPassword,
	}

	auth := socks5.UserPassAuthenticator{
		Credentials: credentials,
	}

	dialer, err := createDial(proxyAddress)
	if err != nil {
		return fmt.Errorf("create dial: %w", err)
	}

	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}

	srv := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(tEditLogger, "", log.LstdFlags))),
		socks5.WithAuthMethods([]socks5.Authenticator{auth}),
		socks5.WithDial(dialContext),
	)

	go func() {
		if err := srv.ListenAndServe("tcp", ":"+serverPort, done); err != nil {
			chErr <- err

			return
		}

		chErr <- errStopServer
	}()

	return nil
}

func createDial(proxyAddress string) (proxy.Dialer, error) {
	u, err := url.Parse(proxyAddress)
	if err != nil {
		return nil, fmt.Errorf("url: parse: %w", err)
	}

	auth := new(proxy.Auth)

	password, ok := u.User.Password()
	if ok {
		auth.User = u.User.Username()
		auth.Password = password
	}

	dialer, err := proxy.SOCKS5("tcp", u.Host, auth, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("proxy: socks5: %w", err)
	}

	return dialer, nil
}

func stopServer(
	pushButtonStartServer *widgets.QPushButton,
	pushButtonStopServer *widgets.QPushButton,
	serverPort string,
	done chan struct{},
	tEditLogger *logger.TEditLogger,
) {
	pushButtonStopServer.ConnectClicked(func(bool) {
		if len(done) == 0 {
			done <- struct{}{}
		}

		pushButtonStartServer.SetEnabled(true)
		pushButtonStopServer.SetEnabled(false)

		if err := stopServerSelfRequest(serverPort, tEditLogger); err != nil {
			var opErr *net.OpError
			if errors.As(err, &opErr) {
				tEditLogger.InsertText(logger.Info, "stop server", "ok")

				return
			}

			tEditLogger.InsertText(logger.Fatal, "stop server: self request", err.Error())
		}
	})
}

func stopServerSelfRequest(serverPort string, tEditLogger *logger.TEditLogger) error {
	conn, err := net.Dial("tcp", "127.0.0.1:"+serverPort)
	if err != nil {
		return fmt.Errorf("net dial: %w", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			tEditLogger.InsertText(logger.Error, "conn close", err.Error())
		}
	}()

	if _, err := conn.Write(nil); err != nil {
		return fmt.Errorf("coon: write: %w", err)
	}

	return nil
}
