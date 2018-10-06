#TODO \
1) Docker build \
BINDIR			:=	$(GOPATH)/bin
BINARYFILE		:=	go-mssql-runner
GOMETALINTER 	:= 	$(BIN_DIR)/gometalinter
#Version is driven by git tags
VERSION			:=	$(shell git describe --tags)
LINUXFOLDER		:=	release/linux
WINFOLDER		:=	release/windows
DARWINFOLDER	:=	release/darwin
ALPINEFOLDER	:=	release/alpine-linux
RELEASEFOLDER	:=	./release
all: test create-folders build zip

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

lint: $(GOMETALINTER)
	gometalinter ./... --vendor

create-folders:
	mkdir -p $(LINUXFOLDER)
	mkdir -p $(WINFOLDER)
	mkdir -p $(DARWINFOLDER)
	mkdir -p $(ALPINEFOLDER)

test: 
	go test -v --cover ./...

build: linux windows darwin alpine-linux

zip:
	zip -r $(RELEASEFOLDER)/$(VERSION)-go-mssql-runner_alpine.zip ./release/alpine-linux
	zip -r $(RELEASEFOLDER)/$(VERSION)-go-mssql-runner_darwin.zip ./release/darwin
	zip -r $(RELEASEFOLDER)/$(VERSION)-go-mssql-runner_linux.zip ./release/linux
	zip -r $(RELEASEFOLDER)/$(VERSION)-go-mssql-runner_windows.zip ./release/windows

linux:
	@echo ------------------Building Linux binary------------------
	GOOS=linux GOARCH=amd64 go build -a -installsuffix -i -v -o $(LINUXFOLDER)/$(BINARYFILE) -ldflags=all='-X "main.Version=$(VERSION)" -s -w'
	@echo ------------------Completed building Linux binary------------------
windows:
	@echo ------------------Building Windows binary------------------
	GOOS=windows GOARCH=amd64 go build -a -installsuffix -i -v -o $(WINFOLDER)/$(BINARYFILE).exe -ldflags=all='-X "main.Version=$(VERSION)" -s -w'
	@echo ------------------Completed building Windows binary------------------
darwin:
	@echo ------------------Building Darwin binary------------------
	GOOS=darwin GOARCH=amd64 go build -a -installsuffix -i -v -o $(DARWINFOLDER)/$(BINARYFILE) -ldflags=all='-X "main.Version=$(VERSION)" -s -w'
	@echo ------------------Completed building Darwin binary------------------
alpine-linux:
	@echo ------------------Building Alpine-Linux binary------------------
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix -v -o $(ALPINEFOLDER)/$(BINARYFILE) -ldflags=all='-X "main.Version=$(VERSION)" -s -w'
	@echo ------------------Completed building Alpine-Linux binary------------------
