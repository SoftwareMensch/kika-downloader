export GOPATH:=${GOPATH}:$(shell pwd)

all:
	GOPATH=${GOPATH} GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o kika-downloader-linux-386 kika-downloader
	GOPATH=${GOPATH} GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o kika-downloader-linux-amd64 kika-downloader

	GOPATH=${GOPATH} GOOS=darwin GOARCH=386 go build -ldflags="-s -w" -o kika-downloader-darwin-386 kika-downloader
	GOPATH=${GOPATH} GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o kika-downloader-darwin-amd64 kika-downloader

	GOPATH=${GOPATH} GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o kika-downloader-windows-386.exe kika-downloader
	GOPATH=${GOPATH} GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o kika-downloader-windows-amd64.exe kika-downloader

