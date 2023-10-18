FROM golang:1.19 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o semver .

FROM alpine:3.14
COPY --from=builder /app/semver /semver
ENTRYPOINT ["/semver"]
