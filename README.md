# GophKeeper

## Как собрать
1) Сборка под Linux
   1) GOOS=linux GOARCH=amd64  go build -o bin/linux64_client cmd/client/main.go
   2) GOOS=linux GOARCH=amd64  go build -o bin/linux64_server cmd/server/main.go
2) Сборка под Windows
   1) GOOS=windows GOARCH=amd64  go build -o bin/win64_client.exe cmd/client/main.go
   2) GOOS=windows GOARCH=amd64  go build -o bin/win64_server.exe cmd/server/main.go
3) Сборка под mac(intel)
   1) GOOS=darwin GOARCH=amd64 go build -o bin/mac64_client cmd/client/main.go
   2) GOOS=darwin GOARCH=amd64 go build -o bin/mac64_server cmd/server/main.go
3) Сборка под mac(silicon)
    1) GOOS=darwin GOARCH=arm64 go build -o bin/mac64_client cmd/client/main.go
    2) GOOS=darwin GOARCH=arm64 go build -o bin/mac64_server cmd/server/main.go

## Как запустить
###  сервер
`./bin/server -g "localhost:3200" -m "./migrations" -f "./files" -d "postgresql://postgres:password@10.66.66.3:5432/postgres?sslmode=disable"`
###  клиент
`./bin/client -g "localhost:3200"` 
адрес сервера можно поменять после запуска клиента если не указали при старте