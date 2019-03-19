#Sudu Backend
-include Makefile.env

DATE = $(shell echo `date +%FT%T%z`)
REGISTRY = "registry.xundupdf.com:5000"
LDFLAGS += -X 'fxlibraries/version._VERSION_=$(VERSION)' -X 'fxlibraries/version._DATE_=$(DATE)'


BUILD_TAGS += gm no_development

.PHONY: all build clean deploy base pdfcenter

all: build

deploy: gocms

test:
	@echo "		Testing gocms..."
	@export GOPATH="$(CURDIR)/../.." && \
	go test -timeout 1m `go list ./... | egrep -v "vendor"`

../fxlibraries:
	@echo "        Cloning fxlibraries..."
	@-cd .. && git clone git@github.com:lvcarrot/fxlibraries.git --branch master || true
	@-[[ ! -e ../fxlibraries ]] && cd .. && git clone git@github.com:lvcarrot/fxlibraries.git || true

build: ../fxlibraries
	@echo "        Building Sudu backend services..."
	@echo "        Version: ${VERSION}"
	@export GOPATH="$(CURDIR)/vendor:$(CURDIR)/../.." && \
        export APP_VERSION="$(VERSION)" && \
        export PKG_CONFIG_PATH="/usr/local/lib/pkgconfig" && \
	go install -v -tags="$(BUILD_TAGS)" -ldflags "$(LDFLAGS)" gocms/...
	@echo "        Done. Binary files are located in GOPATH"

gocms:
	cp -r $(CURDIR)/../../bin/gocms ./
	docker build -t "$(REGISTRY)/sd-gocms:latest" -t "$(REGISTRY)/sd-gocms:$(VERSION)" ./
	docker push $(REGISTRY)/sd-gocms
	rm -fr ./gocms

clean:
	@echo "        Cleaning up..."
	@rm -rf $(CURDIR)/vendor/pkg || true && \
	rm -rf $(CURDIR)/../../pkg || true && \
	rm -f $(CURDIR)/../../built || true
