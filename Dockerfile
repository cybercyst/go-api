FROM golang:1.18-alpine as builder

ARG REVISION

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w \
	-X github.com/cybercyst/go-api/pkg/version.REVISION=${REVISION}" \
	-a -o bin/go-api

FROM alpine

ARG BUILD_DATE
ARG VERSION
ARG REVISION

LABEL maintainer="Forrest Loomis (forrest.loomis@docker.com)"

WORKDIR /home/app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bin/go-api .

ENV GIN_MODE=release
ENTRYPOINT ["./go-api"]

EXPOSE 8080
