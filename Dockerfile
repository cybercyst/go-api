ARG GO_VERSION=1.24

# ---------------------------- base stage -------------------------------
FROM golang:${GO_VERSION} AS base

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ---------------------------- lint stage -------------------------------
FROM base AS lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    golangci-lint run ./...

# ---------------------------- test stage -------------------------------
FROM base AS test
RUN go test ./...

# ---------------------------- build stage -------------------------------
FROM base AS dev
ENTRYPOINT ["go", "run", "."]

# ---------------------------- build stage -------------------------------
FROM base AS build
RUN go build -o /bin/server .

# ---------------------------- prod stage -------------------------------
FROM gcr.io/distroless/static-debian11 as prod

WORKDIR /app
COPY --from=build /bin/server /bin/server

EXPOSE 3000
ENTRYPOINT ["/bin/server"]
