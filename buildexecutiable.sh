rm deploy/linux/rssDedupConfig
GOOS=linux GOARCH=amd64 go build -o deploy/linux/rssDedupConfig cmd/rssdedupconfig/main.go