default: test 

test:            
	go test -v -race ./...

build: 
	go build -v -race -o gitlist

addPath:
	ln -s `pwd`/gitlist /usr/local/bin

install: test build addPath

