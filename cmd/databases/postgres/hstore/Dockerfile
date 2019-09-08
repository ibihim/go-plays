ARG GO_VERSION=1.13

# Build executable.
FROM golang:${GO_VERSION} AS builder
ENV CGO_ENABLED=0

WORKDIR /workspace
COPY ./ ./
RUN go build -o /app .

# Create executable image.
FROM alpine:latest AS final

COPY --from=builder /app /app
RUN  chown -R nobody:nobody /app
USER nobody

ENTRYPOINT ["/app"]

