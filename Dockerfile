FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
