VERSION = 0.0.1

BIN_DIRECTORY   = _bin
EXECUTABLE_NAME = dotnet-launcher
DIST_ZIP_PREFIX = $(EXECUTABLE_NAME).v$(VERSION)

REPO_BASE     = github.com/tintoy
REPO_ROOT     = $(REPO_BASE)/dotnet-launcher

default: fmt build

fmt:
	go fmt $(REPO_ROOT)

clean:
	rm -rf $(BIN_DIRECTORY) $(VERSION_INFO_FILE)
	go clean $(REPO_ROOT)/...

# Peform a development (current-platform-only) build.
dev: fmt
	go build -o $(BIN_DIRECTORY)/$(EXECUTABLE_NAME)

# Perform a full (all-platforms) build.
build: build-windows64 build-linux64 build-mac64

build-windows64:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIRECTORY)/windows-amd64/$(EXECUTABLE_NAME)

build-linux64:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIRECTORY)/linux-amd64/$(EXECUTABLE_NAME)

build-mac64:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIRECTORY)/darwin-amd64/$(EXECUTABLE_NAME)

# Produce archives for a GitHub release.
dist: build
	cd $(BIN_DIRECTORY)/windows-amd64 && \
		zip -9 ../$(DIST_ZIP_PREFIX).windows-amd64.zip $(EXECUTABLE_NAME)
	cd $(BIN_DIRECTORY)/linux-amd64 && \
		zip -9 ../$(DIST_ZIP_PREFIX).linux-amd64.zip $(EXECUTABLE_NAME)
	cd $(BIN_DIRECTORY)/darwin-amd64 && \
		zip -9 ../$(DIST_ZIP_PREFIX)-darwin-amd64.zip $(EXECUTABLE_NAME)
