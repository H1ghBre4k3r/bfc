.DEFAULT_GOAL := build

# run program in dev mode (via go run)
run:
	go run ./cmd/go-bf/main.go

# run all tests
test: 
	go test ./internal/...

# build program for development
build: 
	./scripts/build.sh

# build program for release (without debug information)
release: 
	PESCA_RELEASE=1 ./scripts/build.sh
	
# build program for release (without debug information)
gui: 
	PESCA_RELEASE=1 ./scripts/build.sh

static: 
	STATIC_LINK=1 PESCA_RELEASE=1 ./scripts/build.sh

# clean build folder
clean: 
	rm -rf bin/*
