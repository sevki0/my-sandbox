FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]