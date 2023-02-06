package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
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
	chStop := make(chan struct{}, 0)

	startServer(
		w.PushButtonServerStart,
		w.PushButtonServerStop,
		w.LineEditProxyAddress,
		w.LineEditserverLogin,
		w.LineEditServerPassword,
		w.LineEditServerPort,
		w.TextEditLog,
		chStop,
	)

	stopServer(
		w.PushButtonServerStart,
		w.PushButtonServerStop,
		chStop,
	)
}

func startServer(
	pushButtonStartServer *widgets.QPushButton,
	pushButtonStopServer *widgets.QPushButton,
	lineEditProxyAddress *widgets.QLineEdit,
	lineEditServerLogin *widgets.QLineEdit,
	lineEditServerPassword *widgets.QLineEdit,
	lineEditServerPort *widgets.QLineEdit,
	textEditLog *widgets.QTextEdit,
	chStop chan struct{},
) {
	pushButtonStartServer.ConnectClicked(func(bool) {
		proxyAddress := strings.TrimSpace(lineEditProxyAddress.Text())

		if err := checkProxyAddress(proxyAddress); err != nil {
			textEditLog.InsertPlainText(logger.Logger(logger.Error, err.Error()))

			return
		}

		textEditLog.InsertPlainText(logger.Logger(logger.Info, "start server"))

		pushButtonStartServer.SetEnabled(false)
		pushButtonStopServer.SetEnabled(true)

		if err := startSocks5(
			proxyAddress,
			strings.TrimSpace(lineEditServerLogin.Text()),
			strings.TrimSpace(lineEditServerPassword.Text()),
			strings.TrimSpace(lineEditServerPort.Text()),
			chStop,
		); err != nil {
			textEditLog.InsertPlainText(logger.Logger(logger.Fatal, err.Error()))
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
	chStop chan struct{},
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
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		socks5.WithAuthMethods([]socks5.Authenticator{auth}),
		socks5.WithDial(dialContext),
	)

	chErr := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe("tcp", ":"+serverPort, chStop); err != nil {
			chErr <- err

			return
		}

		chErr <- errStopServer
	}()

	return <-chErr
}

func createDial(proxyAddress string) (proxy.Dialer, error) {
	u, err := url.Parse(proxyAddress)
	if err != nil {
		return nil, fmt.Errorf("url: parse: %w", err)
	}

	password, _ := u.User.Password()

	auth := &proxy.Auth{
		User:     u.User.Username(),
		Password: password,
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
	chStop chan struct{},
) {
	pushButtonStopServer.ConnectClicked(func(bool) {
		pushButtonStartServer.SetEnabled(true)
		pushButtonStopServer.SetEnabled(false)

		chStop <- struct{}{}
	})
}
