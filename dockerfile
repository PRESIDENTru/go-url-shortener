FROM golang:1.25.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o url-shortener ./cmd/server

EXPOSE 8080

CMD ["./url-shortener"]

