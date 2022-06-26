FROM golang:1.18-alpine3.15 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main main.go

FROM golang:1.18-alpine3.15
WORKDIR /app
COPY --from=builder /main .
EXPOSE 8000
ENTRYPOINT [./app/main]