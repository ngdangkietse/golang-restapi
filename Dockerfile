FROM golang:1.20.0-alpine
WORKDIR /go/src/app
COPY . .
RUN go build -o main main.go
CMD ["./main"]
EXPOSE 6969