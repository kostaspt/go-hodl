PKG_LIST = $(shell go list ./...  | grep -v /vendor/)

build:
	@echo "> Building binary..."
	@echo ""
	go build -o ./bin/hodl main.go

test:
	@echo "> Testing..."
	@echo ""
	go test -short ${PKG_LIST}

test-race:
	@echo "> Race Testing..."
	@echo ""
	go test -race ${PKG_LIST}

test-coverage:
	@echo "> Testing with Coverage..."
	@echo ""
	go test -coverprofile=coverage.out ${PKG_LIST}

clean:
	@echo "> Cleaning..."
	@echo ""
	rm -rf ./bin

format:
	@echo "> Formatting..."
	@echo ""
	go fmt ./...