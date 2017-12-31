#TODO \
1) Docker build \
2) Build for alpine linux
BINDIR			:=	$(GOPATH)/bin
BINARYFILE		:=	go-mssql-runner
GOMETALINTER 	:= 	$(BIN_DIR)/gometalinter
#Version is driven by git tags
VERSION			:=	$(shell git describe --tags)
LINUXFOLDER		:=	release/linux
WINFOLDER		:=	release/windows
DARWINFOLDER	:=	release/darwin
ALPINEFOLDER	:=	release/alpine-linux

all: test create-folders build

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

lint: $(GOMETALINTER)
	gometalinter ./... --vendor

create-folders:
	mkdir -p $(LINUXFOLDER)
	mkdir -p $(WINFOLDER)
	mkdir -p $(DARWINFOLDER)

test: update-package
	go test -v --cover ./...

update-package:
	go get -u github.com/denisenkom/go-mssqldb
	go get -u github.com/sirupsen/logrus
	go get -u github.com/spf13/cobra

build: linux windows darwin
	
linux:
	@echo ------------------Building Linux binary------------------
	GOOS=linux GOARCH=amd64 go build -a -installsuffix -i -v -o $(LINUXFOLDER)/$(BINARYFILE) -ldflags="-X main.Version=$(VERSION)"
	@echo ------------------Completed building Linux binary------------------
windows:
	@echo ------------------Building Windows binary------------------
	GOOS=windows GOARCH=amd64 go build -a -installsuffix -i -v -o $(WINFOLDER)/$(BINARYFILE).exe -ldflags="-X main.Version=$(VERSION)"
	@echo ------------------Completed building Windows binary------------------
darwin:
	@echo ------------------Building Darwin binary------------------
	GOOS=darwin GOARCH=amd64 go build -a -installsuffix -i -v -o $(DARWINFOLDER)/$(BINARYFILE) -ldflags="-X main.Version=$(VERSION)"
	@echo ------------------Completed building Darwin binary------------------