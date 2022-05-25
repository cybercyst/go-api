FROM golang:1.18-alpine as builder

ARG REVISION

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w \
	-X github.com/cybercyst/go-api/pkg/version.REVISION=${REVISION}" \
	-a -o bin/go-api

FROM alpine:3.16

ARG BUILD_DATE
ARG VERSION
ARG REVISION

LABEL maintainer="Forrest Loomis (forrest.loomis@docker.com)"

RUN addgroup -S app \
	&& adduser -S -G app app \
	&& apk --no-cache add \
	ca-certificates curl netcat-openbsd

WORKDIR /home/app

COPY --from=builder /app/bin/go-api .

USER app

ENV GIN_MODE=release
ENTRYPOINT ["./go-api"]

