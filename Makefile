# Makefile for building code on different platforms

# Variables
IMAGE_NAME := my-telegram-bot
VERSION := latest

# Targets
linux:
	GOOS=linux GOARCH=amd64 go build -o $(IMAGE_NAME)-linux-amd64 .
	@echo "Linux build complete"

arm:
	GOOS=linux GOARCH=arm go build -o $(IMAGE_NAME)-linux-arm .
	@echo "ARM build complete"

mac:
	GOOS=darwin GOARCH=amd64 go build -o $(IMAGE_NAME)-mac-amd64 .
	@echo "Mac build complete"

windows:
	GOOS=windows GOARCH=amd64 go build -o $(IMAGE_NAME)-windows-amd64.exe .
	@echo "Windows build complete"

clean:
	rm -f $(IMAGE_NAME)-*

.PHONY: linux arm mac windows clean

docker:
	docker build -t $(IMAGE_NAME):$(VERSION) .

clean:
	rm -f $(IMAGE_NAME)-*

.PHONY: linux arm mac windows docker clean
