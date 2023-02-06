package server

import (
	"errors"
	"fmt"
	"regexp"
)

var errWrongValue = errors.New("wrong value")

type RegularExpressions struct {
	ProxySocks5    *regexp.Regexp
	ServerLogin    *regexp.Regexp
	ServerPassword *regexp.Regexp
	ServerPort     *regexp.Regexp
}

func CompileRegexps() RegularExpressions {
	return RegularExpressions{
		ProxySocks5:    regexp.MustCompile(`^socks5://(.+:.+@|)(\d{1,3}[.]){3}\d{1,3}:\d{1,5}$`),
		ServerLogin:    regexp.MustCompile(`^[^\s]{3,}$`),
		ServerPassword: regexp.MustCompile(`^[^\s]{3,}$`),
		ServerPort:     regexp.MustCompile(`^\d{2,5}$`),
	}
}

func checkValues(regexps RegularExpressions, proxyAddress, serverLogin, serverPassword, serverPort string) error {
	if !regexps.ProxySocks5.MatchString(proxyAddress) {
		return fmt.Errorf("proxy address: %s: %v", proxyAddress, errWrongValue)
	}

	if !regexps.ServerLogin.MatchString(serverLogin) {
		return fmt.Errorf("server login: %s: %v", serverLogin, errWrongValue)
	}

	if !regexps.ServerPassword.MatchString(serverPassword) {
		return fmt.Errorf("server password: %s: %v", serverPassword, errWrongValue)
	}

	if !regexps.ServerPort.MatchString(serverPort) {
		return fmt.Errorf("server port: %s: %v", serverPort, errWrongValue)
	}

	return nil
}
