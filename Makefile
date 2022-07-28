# Binary name
BINARY=btcspy
VERSION=0.1
LDFLAGS='-w -s'

build:
	# Build
	rm -f ./${BINARY}
	go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go

mac:
	# Build for mac
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go

mac-arm:
	# Build for mac
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go

linux:
	# Build for linux
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go

linux-arm:
	# Build for linux
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go

windows:
	# Build for windows
	rm -f ./${BINARY}.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe -ldflags ${LDFLAGS} cmd/*.go

windows-arm:
	# Build for windows
	rm -f ./${BINARY}.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o ${BINARY}.exe -ldflags ${LDFLAGS} cmd/*.go

release:
	# Make release for all
	@rm -rf release/ && mkdir release/

	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	@mv ./${BINARY} release/
	@cp -r assets release/

	# zip file
	@zip -r release.zip release

	# clear
	@rm -f ./${BINARY}
	@rm -rf release

clean:
	# Clean projects
	rm -f ${BINARY}
	rm -f ${BINARY}.exe
	rm -f release.zip

.PHONY: build mac mac-arm linux linux-arm windows windows-arm release clean
