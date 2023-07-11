set GOOS=linux
set GOARCH=amd64
go build -o game  .

docker build -t game:v1 .
del game