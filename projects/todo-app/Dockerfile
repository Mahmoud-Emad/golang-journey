FROM golang:latest

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build

CMD [ "go", "run", "main.go" ]

EXPOSE 8080