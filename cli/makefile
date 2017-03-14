export GOPATH:=${GOPATH}

all:
	GOPATH=${GOPATH} GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o kika-downloader-linux-386
	GOPATH=${GOPATH} GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o kika-downloader-linux-amd64

	GOPATH=${GOPATH} GOOS=darwin GOARCH=386 go build -ldflags="-s -w" -o kika-downloader-darwin-386
	GOPATH=${GOPATH} GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o kika-downloader-darwin-amd64

	GOPATH=${GOPATH} GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o kika-downloader-windows-386.exe
	GOPATH=${GOPATH} GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o kika-downloader-windows-amd64.exe

