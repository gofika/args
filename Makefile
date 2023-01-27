LOCAL_DIR := bin
MAIN := args
ROOT := ./cmd/args

.PHONY: linux
linux:
	@echo 'building for linux...'
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "${LOCAL_DIR}/${MAIN}" ${ROOT}

.PHONY: windows
windows:
	@echo 'building for windows...'
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "${LOCAL_DIR}/${MAIN}.exe" ${ROOT}