FROM golang:1.16

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

WORKDIR /app
COPY . /app

RUN go install

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080

CMD ["go", "run", "main.go"]