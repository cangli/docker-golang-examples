FROM golang:latest as builder
WORKDIR /go
RUN git clone https://github.com/cangli/docker-golang-examples.git && \
    unset GOPATH && \
    go mod init main && \
    go get -u github.com/ugorji/go@v1.1.7 && \
    CGO_ENABLED=0 GOOS=linux go build docker-golang-examples/server.go && \
    cp docker-golang-examples/wait-for.sh . && \
    chmod 755 wait-for.sh

FROM alpine:latest
COPY --from=builder /go/server go/wait-for.sh /bin/
CMD ["sh"]
EXPOSE 8080
