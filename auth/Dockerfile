FROM golang
WORKDIR /go/src/ddhub/auth
COPY main.go .
RUN go build main.go
EXPOSE 8080
ENTRYPOINT ["./main"]