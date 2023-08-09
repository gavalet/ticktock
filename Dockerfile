FROM golang:1.18.1 AS builder
RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the webservice
RUN CGO_ENABLED=0 GOOS=linux go build -o myapi cmd/ticktock/main.go

FROM alpine:latest
RUN apk add --no-cache tzdata
WORKDIR /app
# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/myapi .
CMD ["./myapi" ,"172.16.0.2", "8080"]