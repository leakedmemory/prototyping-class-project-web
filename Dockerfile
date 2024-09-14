ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN make build

FROM debian:bookworm

COPY --from=builder /usr/src/app/main /usr/local/bin/
CMD ["/usr/local/bin/main"]
