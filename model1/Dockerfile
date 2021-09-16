FROM golang
WORKDIR /go/src/ddhub/model1
COPY main.go .
RUN go build main.go
EXPOSE 8080
ENTRYPOINT ["./main"]