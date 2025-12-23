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

# Install yt-dlp via pip to get the absolute latest version (nightly) 
# which is required to fix the "Sign in" errors.
RUN python3 -m venv /venv
ENV PATH="/venv/bin:$PATH"
RUN pip install --no-cache-dir -U "yt-dlp[default]"

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
