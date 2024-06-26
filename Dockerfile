ARG GOLANG_VERSION=1.22.0

FROM --platform=$BUILDPLATFORM golang:${GOLANG_VERSION}-alpine AS build

ARG TARGETARCH
ARG TARGETOS

WORKDIR /app
COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /entrypoint cmd/entrypoint/main.go
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /worker cmd/worker/main.go
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /migrate cmd/migrate/main.go


FROM --platform=$BUILDPLATFORM alpine:latest

#LABEL org.opencontainers.image.source = "https://github.com/lognitor/entrypoint"

WORKDIR /app

COPY --from=build /entrypoint /app/entrypoint
COPY --from=build /worker /app/worker

COPY --from=build /migrate /app/migrate
COPY --from=build /app/migrations /app/migrations

CMD ["/app/entrypoint"]