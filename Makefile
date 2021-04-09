BINARY_NAME=vlt
BINARY_PATH=/etc/local/bin/vlt

all:
	make build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$(BINARY_NAME)" -v
	cp ${BINARY_PATH} && chmod +x ${BINARY_PATH}
