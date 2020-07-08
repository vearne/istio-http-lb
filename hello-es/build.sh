export CGO_ENABLED=0
export GOOS=linux
go build -ldflags '-s' hello_es.go
docker build -t vearne/hello-es .
