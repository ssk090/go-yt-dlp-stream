FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy module files
COPY go.mod ./
# COPY go.sum ./

# Download dependencies
# RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Install dependencies: python3 (required by yt-dlp), ffmpeg, and nodejs (for JS execution)
RUN apk add --no-cache python3 ffmpeg curl nodejs

# Install yt-dlp
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
