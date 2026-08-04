FROM golang:1.24-bookworm

LABEL maintainer="tvremote"

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y \
        git \
        make \
        curl \
        wget \
        unzip \
        vim \
        nano \
        tree \
        jq \
        adb \
        ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# ---------------------------------------
# Go Development Tools
# ---------------------------------------

RUN go install golang.org/x/tools/gopls@v0.20.0

RUN go install golang.org/x/tools/cmd/goimports@v0.20.0

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

RUN go install honnef.co/go/tools/cmd/staticcheck@2025.1

RUN go install github.com/go-delve/delve/cmd/dlv@v1.24.2

RUN go install github.com/air-verse/air@v1.61.7

ENV PATH=$PATH:/go/bin

WORKDIR /workspace

CMD ["bash"]