.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t scheduly:v0.1 .

start-container:
	docker run --name scheduly --env-file .env scheduly:v0.1
