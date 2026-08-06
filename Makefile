APP := tvremote
BUILD_DIR := build

ADB_ADDRESS ?= host.docker.internal:16416
TVREMOTE_MODE ?= adb

.PHONY: run build android clean

fmt:
	goimports -w .

lint:
	golangci-lint run

check:
	staticcheck ./...

test:
	go test ./...

run:
	ADB_ADDRESS=$(ADB_ADDRESS) \
	go run ./cmd/tvremote

build:
	go build -o $(BUILD_DIR)/$(APP) ./cmd/tvremote

android:
	CGO_ENABLED=0 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	go build \
		-tags "netgo osusergo" \
		-ldflags="-s -w" \
		-o build/tvremote \
		./cmd/tvremote

clean:
	rm -rf $(BUILD_DIR)