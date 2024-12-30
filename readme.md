# Cara Menjalankan Melalui Terminal

- buka terminal menggunakan bash atau zsh masuk ke dalam project, cd /go-123-api
- go run \*.go
- buka satu lagi terminal untuk melakukan http request melalui curl
- curl http://localhost:9090/users/10
- curl -H "Authorization: Bearer token" http://localhost:9090/users/10
