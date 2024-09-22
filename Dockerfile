ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN make build
RUN mkdir /data

FROM debian:bookworm

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /usr/local/bin/
COPY --from=builder /app/web/static /usr/local/bin/web/static

CMD ["/usr/local/bin/main"]
