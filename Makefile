win32:
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/win/32/bgs.exe
win64:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/win/64/bgs.exe
mac32:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o bin/mac/32/bgs
mac64:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/mac/64/bgs
linux32:
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/linux/32/bgs
linux64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/64/bgs
clear:
	@rm -rf bin/mac/*
	@rm -rf bin/linux/*
	@rm -rf bin/win/*
all: clear win32 win64 mac32 mac64 linux32 linux64