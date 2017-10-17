FROM golang:1.9
WORKDIR /app
ADD gorest /app/
ENTRYPOINT ["./gorest"]
