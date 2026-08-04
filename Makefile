run:
	go run ./cmd/tvremote

build:
	go build \
	-o build/tvremote \
	./cmd/tvremote

android:
	CGO_ENABLED=0 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	go build \
	-o build/tvremote \
	./cmd/tvremote