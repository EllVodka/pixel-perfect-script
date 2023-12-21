BINARY_NAME=scriptPixelPerfect
MAIN_PATH=.
MAIN_FILE=${MAIN_PATH}/main.go
BIN_PATH=${MAIN_PATH}/bin/

build: darwin linux windows ## generate binary for linux, darwin and windows OS

darwin: ## create a make rule for each OS declared on os variable. Now, you can build a binary for linux by running `make linux` command
	go generate ${MAIN_FILE}
	GOARCH=amd64 GOOS=$@ go build -o ${BIN_PATH}${BINARY_NAME}-$@ ${MAIN_FILE}

linux: ## create a make rule for each OS declared on os variable. Now, you can build a binary for linux by running `make linux` command
	go generate ${MAIN_FILE}
	GOARCH=amd64 GOOS=$@ go build -o ${BIN_PATH}${BINARY_NAME}-$@ ${MAIN_FILE}

windows: ## create a make rule for each OS declared on os variable. Now, you can build a binary for linux by running `make linux` command
	go generate ${MAIN_FILE}
	GOARCH=amd64 GOOS=$@ go build -o ${BIN_PATH}${BINARY_NAME}-$@.exe ${MAIN_PATH}

.PHONY := run
run:
	go run ${MAIN_FILE} 

.PHONY: dev
dev: build run ## build and run application in dev mode

.PHONY: clean_build
clean_build:
	go clean
	rm ${BIN_PATH}${BINARY_NAME}-darwin
	rm ${BIN_PATH}${BINARY_NAME}-linux
	rm ${BIN_PATH}${BINARY_NAME}-windows.exe
