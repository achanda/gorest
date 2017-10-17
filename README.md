Basic REST API in Go

go run main.go

Create a new post
curl -H "Content-Type: application/json" -X POST -d '{"title": "Hello world", "text": "intro"}' http://127.0.0.1:3000/posts

Get posts
curl http://127.0.0.1:3000/posts

Build
GOOS=linux GOARCH=amd64 go build -o gorest .

Generate image
docker build . -t gorest
