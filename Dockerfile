FROM golang:1.23.4-alpine as builder
WORKDIR /go/src/
COPY . .
RUN go mod download
RUN go build -o dokjasijeom .
EXPOSE 8080
CMD ["./dokjasijeom"]