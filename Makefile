staticsleep: main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o $@
