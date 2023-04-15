BIN_DIR = ./builds/
BIN_NAME = grade-tracker
SRC = grade-tracker.go

all: linux-amd64 linux-arm64 windows-amd64 windows-arm64 mac-amd64 mac-arm64

linux: linux-amd64 linux-arm64

windows: windows-amd64 windows-arm64

mac: mac-amd64 mac-arm64

linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)$(BIN_NAME)-linux-amd64 $(SRC)

linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o $(BIN_DIR)$(BIN_NAME)-linux-arm64 $(SRC)

windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)$(BIN_NAME)-win-amd64.exe $(SRC)

windows-arm64:
	env GOOS=windows GOARCH=arm64 go build -o $(BIN_DIR)$(BIN_NAME)-win-arm64.exe $(SRC)

mac-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)$(BIN_NAME)-mac-amd64 $(SRC)

mac-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)$(BIN_NAME)-mac-arm64 $(SRC)

