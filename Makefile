OUT=dist
GO=go

all: clean build

build: main.go
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o $(OUT)/envfile

clean:
	rm -f $(OUT)/*