FROM golang:1.20.4

WORKDIR /app

COPY ./src /app

CMD ["go", "run", "main.go"]
