APPNAME := GophKeeper
PACKAGE := main

REVISION := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell date +%F)
#mac_silicon
build_mac_silicon_client:
	GOOS=darwin GOARCH=arm64 go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APPNAME)_client_mac  cmd/client/main.go

run_mac_silicon_client: build_mac_silicon_client
	./bin/$(APPNAME)_client_mac -g "localhost:3200"

build_mac_silicon_server:
	GOOS=darwin GOARCH=arm64 go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APPNAME)_server_mac  cmd/server/main.go

run_mac_silicon_server: build_mac_silicon_server
	./bin/$(APPNAME)_server_mac -g "localhost:3200"
#linux
build_linux_client:
	GOOS=linux GOARCH=amd64 go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APPNAME)_client_linux  cmd/client/main.go

run_linux_client: build_linux_client
	./bin/$(APPNAME)_client_linux -g "localhost:3200"

build_linux_server:
	GOOS=linux GOARCH=amd64 go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APPNAME)_server_linux  cmd/server/main.go

run_linux_server: build_linux_server
	./bin/$(APPNAME)_server_linux -g "localhost:3200"
#windows
build_windows_client:
	GOOS=windows GOARCH=amd64 go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APPNAME)_client_windows.exe  cmd/client/main.go

run_windows_client: build_windows_client
	./bin/$(APPNAME)_client_windows -g "localhost:3200"

build_windows_server:
	GOOS=windows GOARCH=amd64 go build  -ldflags "-X $(PACKAGE).buildCommit=$(REVISION) -X $(PACKAGE).buildVersion=$(BRANCH) -X $(PACKAGE).buildDate=$(DATE)" -o bin/$(APPNAME)_server_windows.exe  cmd/server/main.go

run_windows_server: build_windows_server
	./bin/$(APPNAME)_server_windows -g "localhost:3200"