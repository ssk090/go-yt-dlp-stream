# go-yt-dlp-stream
a simple server to stream youtube audio

## To run this from postman
```go
curl --location 'http://localhost:8080/stream' \
--header 'Content-Type: application/json' \
--data '{
    "title": "vibex sugar"
}'
```

## To run this from terminal
```go
curl -X POST "http://localhost:8080/stream" \
     -H "Content-Type: application/json" \
     -d '{"title": "vibex sugar"}' \
     --silent | ffplay -autoexit -nodisp -
```
