APP := tvremote
BUILD_DIR := build

ADB_ADDRESS ?= 127.0.0.1:5557
TVREMOTE_MODE ?= adb

.PHONY: run build android clean

ANDROID_API := 21

CC_ARM := armv7a-linux-androideabi$(ANDROID_API)-clang
CXX_ARM := armv7a-linux-androideabi$(ANDROID_API)-clang++

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
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	CC=$(CC_ARM) \
	CXX=$(CXX_ARM) \
	go build \
		-o build/tvremote \
		./cmd/tvremote

clean:
	rm -rf $(BUILD_DIR)