# GophKeeper

## Готовые собранные приложения в папке BIN 

## Как собрать
`make build_server`
`make build_client`
## Как запустить что собрали
###  сервер
`./bin/server -g "localhost:3200" -f "./GophKeeperServerFiles" -d "postgresql://postgres:password@10.66.66.3:5432/postgres?sslmode=disable"`
###  клиент
`./bin/client -g "localhost:3200"` 
адрес сервера можно поменять после запуска клиента если не указали при старте