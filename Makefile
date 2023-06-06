VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "0.0.1")
LDFLAGS := -ldflags "-X main.Version=${VERSION}"

run:
	go run ${LDFLAGS} cmd/server/main.go