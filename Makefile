GO=go
NAME := heatman
VERSION := 1.0.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)'
	-X 'main.revision=$(REVISION)'

all: test build

deps:

setup: deps update_version
#	git submodule update --init

update_version:
	@for i in README.md ; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done
	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' cmd/heatman/main.go > a
	@mv a cmd/heatman/main.go
	@echo "Replace version to \"${VERSION}\""

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

build: setup
	cd cmd/heatman; $(GO) build -o $(NAME) -v

clean:
	$(GO) clean
	rm -rf $(NAME)
