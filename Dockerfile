ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

WORKDIR /app
COPY --from=builder /usr/src/app/"Blog Markdowns" ./"Blog Markdowns"
COPY --from=builder /usr/src/app/web/static ./web/static
COPY --from=builder /run-app .
CMD ["./run-app"]
