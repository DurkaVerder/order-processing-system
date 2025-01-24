FROM golang:1.23.3

WORKDIR /APIGateway

COPY . .

RUN go mod tidy

RUN go get -u "github.com/lib/pq"

RUN go build -o app ./cmd/app/main.go

CMD ["./app"]