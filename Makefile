NAME ?= mottainai-server
PACKAGE_NAME ?= $(NAME)
GOLANG_VERSION=$(shell go env GOVERSION)

override LDFLAGS += -X "github.com/MottainaiCI/mottainai-server/pkg/settings.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
override LDFLAGS += -X "github.com/MottainaiCI/mottainai-server/pkg/settings.BuildCommit=$(shell git rev-parse HEAD)"
override LDFLAGS += -X "github.com/MottainaiCI/mottainai-server/pkg/settings.BuildGoVersion=$(GOLANG_VERSION)"


REVISION := $(shell git rev-parse --short HEAD || echo unknown)
VERSION := $(shell git describe --tags || cat pkg/settings/settings.go | echo $(REVISION) || echo dev)
VERSION := $(shell echo $(VERSION) | sed -e 's/^v//g')
BUILD_PLATFORMS ?= -osarch="linux/amd64" -osarch="linux/386" -osarch="linux/arm" -osarch="linux/arm64"

SUBDIRS =
DESTDIR =
UBINDIR ?= /usr/bin
LIBDIR ?= /usr/lib
SBINDIR ?= /sbin
USBINDIR ?= /usr/sbin
BINDIR ?= /bin
LIBEXECDIR ?= /usr/libexec
SYSCONFDIR ?= /etc
LOCKDIR ?= /var/lock
LIBDIR ?= /var/lib
EXTENSIONS ?=
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: deps build

build-test: test multiarch-build

help:
	# make all => deps test lint build
	# make deps - install all dependencies
	# make test - run project tests
	# make lint - check project code style
	# make build - build project for all supported OSes

clean:
	rm -rf release/

deps:
	go env
	# Installing dependencies...
	GO111MODULE=off go get golang.org/x/lint/golint
	GO111MODULE=off go get github.com/mitchellh/gox
	GO111MODULE=off go get golang.org/x/tools/cmd/cover
	GO111MODULE=on go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	GO111MODULE=off go get github.com/onsi/gomega/...
	GO111MODULE=off go get -u github.com/maxbrunsfeld/counterfeiter
	ginkgo version

build-exporter:
ifeq ($(EXTENSIONS),)
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o ./mottainai-exporter/mottainai-exporter ./mottainai-exporter
else
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -tags $(EXTENSIONS) -o ./mottainai-exporter/mottainai-exporter ./mottainai-exporter
endif

build-importer:
ifeq ($(EXTENSIONS),)
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o ./mottainai-importer/mottainai-importer ./mottainai-importer
else
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -tags $(EXTENSIONS) -o ./mottainai-importer/mottainai-importer ./mottainai-importer
endif

build-agent:
ifeq ($(EXTENSIONS),)
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o ./mottainai-agent/mottainai-agent ./mottainai-agent
else
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -tags $(EXTENSIONS) -o ./mottainai-agent/mottainai-agent ./mottainai-agent
endif

build-scheduler:
ifeq ($(EXTENSIONS),)
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o ./mottainai-scheduler/mottainai-scheduler ./mottainai-scheduler
else
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -tags $(EXTENSIONS) -o ./mottainai-scheduler/mottainai-scheduler ./mottainai-scheduler
endif

build-cli:
ifeq ($(EXTENSIONS),)
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o ./mottainai-cli/mottainai-cli ./mottainai-cli
else
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -tags $(EXTENSIONS) -o ./mottainai-cli/mottainai-cli ./mottainai-cli
endif

build-server:
ifeq ($(EXTENSIONS),)
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)'
else
		CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -tags $(EXTENSIONS)
endif

build: build-server build-cli build-agent build-scheduler build-exporter build-importer

lint:
	golint ./... | grep -v "be unexported"

test:
	GO111MODULE=off go get github.com/onsi/ginkgo/v2/ginkgo
	GO111MODULE=off go get github.com/onsi/gomega/...
	ginkgo -r -flake-attempts 3 ./...

.PHONY: test-coverage
test-coverage:
	scripts/ginkgo.coverage.sh --codecov

.PHONY: coverage
coverage:
	go test ./... -coverprofile=coverage.txt -race -covermode=atomic

docker-test:
	docker run -v $(ROOT_DIR)/:/test \
	-e ACCEPT_LICENSE=* \
	--entrypoint /bin/bash -ti --user root --rm mottainaici/test -c \
	"mkdir -p /root/go/src/github.com/MottainaiCI && \
	cp -rf /test /root/go/src/github.com/MottainaiCI/mottainai-server && \
	cd /root/go/src/github.com/MottainaiCI/mottainai-server && \
	make deps test"

compose-test-run: build
		@tmpdir=`mktemp --tmpdir -d`; \
		cp -rf $(ROOT_DIR)/contrib/docker-compose "$$tmpdir"; \
		pushd "$$tmpdir/docker-compose"; \
		mv docker-compose.arangodb.yml docker-compose.yml; \
		trap 'docker-compose down -v --remove-orphans;rm -rf "$$tmpdir"' EXIT; \
		echo ">> Server will be avilable at: http://127.0.0.1:4545" ; \
		sed -i "s|#- ./mottainai-server.yaml:/etc/mottainai/mottainai-server.yaml|- "$(ROOT_DIR)"/mottainai-server:/usr/bin/mottainai-server|g" docker-compose.yml; \
		sed -i "s|# For static config:|- "$(ROOT_DIR)":/var/lib/mottainai|g" docker-compose.yml; \
		docker-compose up

kubernetes:
	make/kubernetes

helm-gen:
	make/helm-gen

install:
	install -d $(DESTDIR)$(LOCKDIR)
	install -d $(DESTDIR)$(BINDIR)
	install -d $(DESTDIR)$(UBINDIR)
	install -d $(DESTDIR)$(SYSCONFDIR)
	install -d $(DESTDIR)$(LIBDIR)

	install -d $(DESTDIR)$(LOCKDIR)/mottainai
	install -d $(DESTDIR)$(SYSCONFDIR)/mottainai
	install -d $(DESTDIR)$(LIBDIR)/mottainai

	install -m 0755 $(NAME) $(DESTDIR)$(UBINDIR)/
	cp -rf templates/ $(DESTDIR)$(LIBDIR)/mottainai
	cp -rf public/ $(DESTDIR)$(LIBDIR)/mottainai

	install -m 0755 contrib/config/mottainai-server.yaml.example $(DESTDIR)$(SYSCONFDIR)/mottainai/

gen-fakes:
	counterfeiter -o tests/fakes/http_client.go pkg/client/client.go HttpClient

.PHONY: multiarch-build
multiarch-build:
	CGO_ENABLED=0 gox $(BUILD_PLATFORMS) -ldflags '$(LDFLAGS)' -output="release/$(NAME)-$(VERSION)-{{.OS}}-{{.Arch}}"

.PHONY: goreleaser-snapshot
goreleaser-snapshot:
	rm -rf dist/ || true
	GOVERSION=$(GOLANG_VERSION) goreleaser release --debug --skip-publish  --skip-validate --snapshot
