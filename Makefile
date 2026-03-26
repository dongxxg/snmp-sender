.PHONY: all serverless deps docker docker-cgo clean docs test test-race test-integration fmt lint install deploy-docs build
DEST_DIR           = ./dist

CMD ?= snmp-sender
importpath=unitechs.com/unios-dice/uni-base/core/startup
VERSION = $(shell cat VERSION 2>/dev/null || echo '0.0.0')
VER_CUT   := $(shell echo $(VERSION) | cut -c2-)
GITREV = $(shell git rev-parse --short HEAD || echo unknown)
GITBRANCH = $(shell git branch | sed -n -e 's/^\* \(.*\)/\1/p')
BUILDTIME = $(shell date +'%Y-%m-%d_%T')
LDFLAGS = -ldflags "-X $(importpath).version=${VERSION} -X $(importpath).gitBranch=${GITBRANCH} -X $(importpath).gitCommit=${GITREV} -X $(importpath).builtAt=${BUILDTIME}"
platform ?= amd64

gotool:
	#go mod tidy
	go fmt ./...
	go vet ./...


all: clean gotool $(TARGET) $(CMD)

deps:
	@go mod tidy

clean:
	@rm -rf ./dist
	@rm -rf ./build
	@rm -rf ${CMD}.tar.gz

$(CMD):
	CGO_ENABLED=0 GOOS=linux GOARCH=$(platform)  go build ${LDFLAGS} -o build/$@/$@ .

pipeline-pack: all
	@if [ -e dist ] ; then rm -rf dist; fi
	@mkdir dist
	@$(foreach var,$(CMD),mkdir -p ./build/$(var)/conf;)
	@$(foreach var,$(CMD),cp -r ./conf/ ./build/$(var)/;)
	@tar -C build -zcf ${CMD}.tar.gz .
	@cp Dockerfile_$(platform) ./dist/Dockerfile
	sed -i 's/{{project}}/$(CMD)/g' ./dist/Dockerfile
	@mv ${CMD}.tar.gz ./dist

.PHONY: p
p: pipeline-pack

test:
	export UNIBASE_CONFIG_PATH=$(CURDIR)/pkg/cmd/conf/configuration.toml
	export PROCESSOR_CONFIG_PATH=$(CURDIR)/pkg/cmd/conf/processor_engine.yaml
	export PROCESSOR_STREAMS_PATH=$(CURDIR)/pkg/cmd/conf/streams
	@go test ./...
t:
	export UNIBASE_CONFIG_PATH=$(CURDIR)/pkg/cmd/conf/configuration.toml
	export PROCESSOR_CONFIG_PATH=$(CURDIR)/pkg/cmd/conf/processor_engine.yaml
	export PROCESSOR_STREAMS_PATH=$(CURDIR)/pkg/cmd/conf/streams
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out