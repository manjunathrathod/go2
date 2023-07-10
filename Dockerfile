# Start from a Debian-based image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.17 AS builder

# Copy the local package files to the container's workspace.
WORKDIR /go/src/app
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/myapp

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/app .
EXPOSE 8080
CMD ["./app"]
