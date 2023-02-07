package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"

	"github.com/things-go/go-socks5"
	"golang.org/x/net/proxy"

	"qt-proxyserver/internal/logger"
	"qt-proxyserver/uigen"
)

var errStopServer = errors.New("stop server")

func Server(w *uigen.UIWindow, regexps RegularExpressions, tEditLogger *logger.TEditLogger) {
	startServer(w, regexps, tEditLogger)
}

func startServer(w *uigen.UIWindow, regexps RegularExpressions, tEditLogger *logger.TEditLogger) {
	chErr := make(chan error, 1)

	go func() {
		for {
			err := <-chErr

			tEditLogger.InsertText(logger.Fatal, "global error", err.Error())
		}
	}()

	w.PushButtonServerStart.ConnectClicked(func(bool) {
		proxyAddress := strings.TrimSpace(w.LineEditProxyAddress.Text())
		serverLogin := strings.TrimSpace(w.LineEditServerLogin.Text())
		serverPassword := strings.TrimSpace(w.LineEditServerPassword.Text())
		serverPort := strings.TrimSpace(w.LineEditServerPort.Text())

		if err := checkValues(regexps, proxyAddress, serverLogin, serverPassword, serverPort); err != nil {
			tEditLogger.InsertText(logger.Error, "check values", err.Error())

			return
		}

		w.PushButtonServerStart.SetEnabled(false)

		tEditLogger.InsertText(logger.Info, "start server", serverPort)

		if err := startSocks5(
			proxyAddress,
			serverLogin,
			serverPassword,
			serverPort,
			chErr,
			tEditLogger,
		); err != nil {
			tEditLogger.InsertText(logger.Fatal, "start socks5", err.Error())
		}
	})
}

func startSocks5(
	proxyAddress string,
	serverLogin string,
	serverPassword string,
	serverPort string,
	chErr chan error,
	tEditLogger *logger.TEditLogger,
) error {
	credentials := socks5.StaticCredentials{
		serverLogin: serverPassword,
	}

	auth := socks5.UserPassAuthenticator{
		Credentials: credentials,
	}

	dialer, err := createDialer(proxyAddress)
	if err != nil {
		return fmt.Errorf("create dial: %w", err)
	}

	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}

	var srv *socks5.Server

	if serverLogin == "" {
		srv = socks5.NewServer(
			socks5.WithDial(dialContext),
			socks5.WithLogger(socks5.NewLogger(log.New(tEditLogger, "", log.LstdFlags))),
		)
	} else {
		srv = socks5.NewServer(
			socks5.WithDial(dialContext),
			socks5.WithAuthMethods([]socks5.Authenticator{auth}),
			socks5.WithLogger(socks5.NewLogger(log.New(tEditLogger, "", log.LstdFlags))),
		)
	}

	go func() {
		if err := srv.ListenAndServe("tcp", ":"+serverPort); err != nil {
			chErr <- fmt.Errorf("listen and serve: %w", err)

			return
		}

		chErr <- errStopServer
	}()

	return nil
}

func createDialer(proxyAddress string) (proxy.Dialer, error) {
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
