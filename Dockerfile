FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go main_test.go ./
RUN go build -o web-app main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/web-app .
COPY static ./static
EXPOSE 8080
CMD ["./web-app"]
