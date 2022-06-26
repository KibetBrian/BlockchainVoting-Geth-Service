FROM golang:1.18-alpine3.15
WORKDIR /app
COPY . .
RUN go mod download
RUN apk add build-base
RUN GOOS=linux go build -o /main
CMD ["/main"]