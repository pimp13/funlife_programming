ARTIFACT_NAME := currency

.PHONY: run build

build:
	@go build -o ./bin/${ARTIFACT_NAME} ./cmd/${ARTIFACT_NAME}/main.go
	@echo compile cli app is ok!

run: build
	@./bin/${ARTIFACT_NAME}