FROM golang:1.20.4

WORKDIR /app

COPY ./src /app

RUN go mod download

CMD ["go", "run", "main.go"]
