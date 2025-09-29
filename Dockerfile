FROM go:1.23-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./app"]
