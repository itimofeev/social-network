FROM golang:1.21-bookworm as builder

# Create and change to the app directory.
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/pressly/goose/v3/cmd/goose@v3.10.0

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o social-network ./cmd/social-network && \
    chmod +x social-network
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o dialogs ./cmd/dialogs && \
    chmod +x dialogs

FROM alpine:3.18

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/bin/goose /goose
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/social-network /app/social-network
COPY --from=builder /app/dialogs /app/dialogs
