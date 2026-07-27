FROM golang:1.26.5-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates && adduser -D -g '' appuser

COPY --from=builder /out/server /app/server

EXPOSE 8080

USER appuser

ENTRYPOINT ["/app/server"]
