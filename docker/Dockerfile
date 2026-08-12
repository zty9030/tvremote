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

# ------------------------------------------------------------
# Android NDK
# ------------------------------------------------------------

ARG NDK_VERSION=r26d

ENV ANDROID_NDK=/opt/android-ndk

ENV PATH=/go/bin:${ANDROID_NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin:${PATH}
# ------------------------------------------------------------
# Android NDK
# This is placed AFTER Go tools because we may change or debug
# the NDK configuration during development.
# ------------------------------------------------------------

RUN mkdir -p /opt && \
    cd /opt && \
    wget -q https://dl.google.com/android/repository/android-ndk-${NDK_VERSION}-linux.zip && \
    unzip -q android-ndk-${NDK_VERSION}-linux.zip && \
    mv android-ndk-${NDK_VERSION} android-ndk && \
    rm android-ndk-${NDK_VERSION}-linux.zip

# ------------------------------------------------------------
# Verify Toolchain
# ------------------------------------------------------------

RUN go version

RUN armv7a-linux-androideabi21-clang --version

WORKDIR /workspace

CMD ["bash"]