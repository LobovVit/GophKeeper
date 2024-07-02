# GophKeeper

## Как собрать

## Как запустить
###  сервер
`./bin/server -g "localhost:3200" -m "./migrations" -f "./files" -d "postgresql://postgres:password@10.66.66.3:5432/postgres?sslmode=disable"`
###  клиент
`./bin/client -g "localhost:3200"` 
адрес сервера можно поменять после запуска клиента если не указали при старте