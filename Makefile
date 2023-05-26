build: 
	go build .

dev:
	nodemon -e go --watch './**/*.go' --signal SIGTERM --exec 'go' run .
.PHONY: dev build
