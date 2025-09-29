FROM golang:1.25

WORKDIR /app

COPY go.mod ./

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./app"]
