FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod init gitservice
RUN go mod tidy
RUN go build -o gitservice .

CMD ["./gitservice"]