qt:
	goqtuic -ui-file main.ui -go-test-file main.go

build:
	go build -o proxyserver main.go