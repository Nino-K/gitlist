default: test 

test:            
	godep restore
	go test -v -race ./...

build: 
	go build -v -race -o gitlist

addPath:
	ln -sf `pwd`/gitlist /usr/local/bin

install: test build addPath

