APP := GophKeeper
PACKAGE := main

REVISION := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell date +%F)

build_client:
	go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APP)_client  cmd/client/main.go

run_client: build
	./bin/$(APPNAME)_client -g "localhost:3200"

build_server:
	go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APP)_server  cmd/server/main.go

run_server: build_server
	./bin/$(APPNAME)_server -g "localhost:3200" -d "postgresql://postgres:password@10.66.66.3:5432/postgres?sslmode=disable"
