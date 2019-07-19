FROM golang:latest as builder
WORKDIR /go
COPY server.go .
RUN unset GOPATH && go mod init main && go get -u github.com/ugorji/go@v1.1.7 && CGO_ENABLED=0 GOOS=linux go build server.go

FROM alpine:latest
COPY --from=builder /go/server /bin/
COPY wait-for.sh /bin/ 
CMD ["sh"]
EXPOSE 8080
