FROM golang:1.22-alpine as build

WORKDIR /app
COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint cmd/main.go

FROM alpine:latest

LABEL org.opencontainers.image.source = "https://github.com/lognitor/entrypoint"

WORKDIR /app

COPY --from=build /entrypoint /app/entrypoint

CMD ["/entrypoint"]