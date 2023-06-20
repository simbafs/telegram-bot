build: 
	go build .

dev:
	nodemon -e go --watch './**/*.go' --signal SIGTERM --exec 'go' run .

image: Dockerfile
	docker build . -t simbafs/bot
.PHONY: dev build image
