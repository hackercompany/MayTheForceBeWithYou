default:fmt
	$(MAKE) build

build: test main_application

main_application:
	go build

test:
	go test ./...

fmt:
	go fmt ./src/...
