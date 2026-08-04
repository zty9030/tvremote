FROM golang:1.24-bookworm

RUN apt-get update && \
    apt-get install -y \
    adb \
    git \
    make \
    ca-certificates

WORKDIR /workspace

CMD ["bash"]