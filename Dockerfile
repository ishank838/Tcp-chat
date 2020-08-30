FROM golang:1.13-alpine3.12

WORKDIR /go/src/

COPY . .

RUN go get ./...

RUN go build -o main .

EXPOSE 8888

CMD ["./main"]