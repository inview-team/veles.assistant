FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

COPY vendor ./vendor

COPY . .

RUN go build -mod=vendor -o veles_assistant main.go

CMD ["./veles_assistant"]
