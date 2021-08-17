CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fabric-agent main.go
# scp -r fabric-agent root@172.38.50.211:/root