FROM golang:1.11

RUN apt-get update && \
    apt-get install -y \
        git make cross-gcc-dev && \
    dpkg --add-architecture armhf && \
    dpkg --add-architecture armel && \
    apt-get install -y \
        crossbuild-essential-armel crossbuild-essential-armhf && \
    rm -rf /var/lib/apt/lists/**

RUN curl -O -L https://github.com/gotestyourself/gotestsum/releases/download/v0.3.2/gotestsum_0.3.2_linux_amd64.tar.gz && \
    tar xf gotestsum_0.3.2_linux_amd64.tar.gz && \
    mv gotestsum /usr/local/bin
RUN go get -u github.com/derekparker/delve/cmd/dlv

WORKDIR /go/src/github.com/rjansen/migi
ENTRYPOINT ["make"]
