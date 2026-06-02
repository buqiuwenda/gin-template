FROM golang:1.24.13-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/gin-template ./cmd

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /bin/gin-template .
COPY configs/config.example.yaml ./configs/config.yaml

EXPOSE 8080
CMD ["./gin-template", "web", "-c", "configs/config.yaml"]
