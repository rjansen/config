FROM golang:1.11

RUN curl -O -L https://github.com/gotestyourself/gotestsum/releases/download/v0.3.2/gotestsum_0.3.2_linux_amd64.tar.gz && \
    tar xf gotestsum_0.3.2_linux_amd64.tar.gz && \
    mv gotestsum /usr/local/bin
RUN go get -u github.com/derekparker/delve/cmd/dlv
RUN useradd -u 1000 -ms /bin/bash migi
USER migi

WORKDIR /go/src/github.com/rjansen/migi
ENTRYPOINT ["make"]
