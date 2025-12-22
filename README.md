# go-yt-dlp-stream
a simple server to stream youtube audio

```go
curl --location 'http://localhost:8080/stream' \
--header 'Content-Type: application/json' \
--data '{
    "title": "vibex sugar"
}'
