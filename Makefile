default: test 

test:            
	dep ensure
	go test -v -race ./...

build: 
	go build -v -race -o gitlist

addPath:
	mv `pwd`/gitlist "${GOPATH}/bin"

install: test build addPath

