FROM golang:latest
WORKDIR /booksalon-go
COPY . .

EXPOSE 8080

RUN go build .
