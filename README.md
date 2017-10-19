Basic REST API in Go

# Running locally (needs posgres running)
```
# set env vars
go run main.go
```

# Create a new post
```
curl -H "Content-Type: application/json" -X POST -d '{"title": "Hello world", "text": "intro"}' http://127.0.0.1:3000/posts
```

# Get version
```
curl http://127.0.0.1:3000/version
```

# Get posts
```
curl http://127.0.0.1:3000/posts
```

# Build
```
GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/achanda/gorest/version.Version=`git rev-parse HEAD`" -o gorest .
```

# Generate image
```
docker build . -t gorest
```
