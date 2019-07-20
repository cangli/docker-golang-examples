FROM golang:latest as builder
WORKDIR /go
RUN git clone https://github.com/cangli/docker-golang-examples.git && \
    unset GOPATH && \
    mv docker-golang-examples/go.mod . && \
    CGO_ENABLED=0 GOOS=linux go build docker-golang-examples/server.go && \
    git clone https://github.com/mrako/wait-for.git && \
    cp wait-for/wait-for wait-for.sh && \
    chmod 755 wait-for.sh

FROM alpine:latest
COPY --from=builder /go/server go/wait-for.sh /bin/
CMD ["sh"]
EXPOSE 8080
