# This is a self-documenting Makefile, see https://www.thapaliya.com/en/writings/well-documented-makefiles/
default: help

GRAFANA_PLUGIN_ID := $(shell jq -r '.id' src/plugin.json)
GRAFANA_PLUGIN_VERSION := $(shell jq -r '.version' package.json)
GRAFANA_PLUGIN_ARTIFACT := $(GRAFANA_PLUGIN_ID)-grafana-$(GRAFANA_PLUGIN_VERSION).zip
GRAFANA_PLUGIN_ARTIFACT_CHECKSUM := $(GRAFANA_PLUGIN_ARTIFACT).md5
RMF_QUERY_PARSER_PATH := ./src/datasources/rmf-datasource/parser/lib
RMF_QUERY_PARSER_FILES := \
	$(RMF_QUERY_PARSER_PATH)/RMFQueryLexer.ts \
	$(RMF_QUERY_PARSER_PATH)/RMFQueryListener.ts \
	$(RMF_QUERY_PARSER_PATH)/RMFQueryParser.ts

all: clean build zip ## Executes clean build and zip tasks

##@ Dependencies

node_modules: package.json yarn.lock
	yarn install --immutable

deps-frontend: node_modules ## Install Node.js dependencies

deps-backend: ## Install Go dependencies
	go get ./pkg

deps: deps-frontend deps-backend ## Install all dependencies

##@ Development

dev-frontend: deps-frontend ## Build frontend data sources (development)
	yarn run dev

watch-frontend: deps-frontend ## Auto rebuilt frontend on file changes
	yarn run watch

dev-backend: deps-backend ## Build backend data sources (development)
	go build -race -o ./dist/datasources/redisrmf-datasource/gpx_ibm-rmf-datasource_$$(go env GOOS)_$$(go env GOARCH) -tags netgo -ldflags -w ./pkg

restart-backend: ## Rebuild and restart backend data source (as root)
	sudo -u "$$(stat -c '%U' .)" make dev-backend
	killall gpx_ibm-rmf-datasource_$$(go env GOOS)_$$(go env GOARCH)

##@ Antlr4
ANTLR_JAR=antlr-4.13.2-complete.jar
ANTLR_JAR_PATH=/usr/local/lib/$(ANTLR_JAR)

.antlr/$(ANTLR_JAR):
	mkdir -p .antlr
ifneq ("$(wildcard $(ANTLR_JAR_PATH))","")
	@echo $(ANTLR_JAR_PATH) found, copying...
	cp $(ANTLR_JAR_PATH) .antlr/$(ANTLR_JAR)
else
	@echo $(ANTLR_JAR_PATH) not found, downloading...
	curl -o .antlr/$(ANTLR_JAR) https://www.antlr.org/download/$(ANTLR_JAR)
endif

##@ Build

$(RMF_QUERY_PARSER_FILES): .antlr/$(ANTLR_JAR) ./grammar/RMFQuery.g4 ## Generate Query Parser
	java -Xmx500M -cp .antlr/$(ANTLR_JAR) org.antlr.v4.Tool -Dlanguage=TypeScript -o $(RMF_QUERY_PARSER_PATH) -Xexact-output-dir ./grammar/RMFQuery.g4

build-frontend: deps-frontend $(RMF_QUERY_PARSER_FILES) ## Build frontend data sources
	yarn build

GO_LD_FLAGS := -w -s -extldflags "-static"
build-backend: deps-backend ## Build backend data source
	for arch in amd64 arm64 s390x ppc64le; do \
	  CGO_ENABLED=0 GOOS=linux GOARCH=$${arch} go build -o ./dist/datasources/rmf-datasource/gpx_ibm-rmf-datasource_linux_$${arch} -ldflags '$(GO_LD_FLAGS)' ./pkg; \
	done
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/datasources/rmf-datasource/gpx_ibm-rmf-datasource_darwin_amd64 -ldflags '$(GO_LD_FLAGS)' ./pkg
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./dist/datasources/rmf-datasource/gpx_ibm-rmf-datasource_darwin_arm64 -ldflags '$(GO_LD_FLAGS)' ./pkg
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/datasources/rmf-datasource/gpx_ibm-rmf-datasource_windows_amd64.exe -ldflags '$(GO_LD_FLAGS)' ./pkg

build-backend-amd64-linux: deps-backend ## Build backend data source
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/datasources/rmf-datasource/gpx_ibm-rmf-datasource_linux_amd64 -ldflags '$(GO_LD_FLAGS)' ./pkg; \

test: deps-backend ## Run unittests
	go test -v ./...

lint:  ## Run linters
	golangci-lint run

build: build-frontend build-backend ## Build everything

zip: ## Zip the plugin
	rm -rf build && mkdir build
	cp -a dist "build/$(GRAFANA_PLUGIN_ID)"
	cd build && zip -r "$(GRAFANA_PLUGIN_ARTIFACT)" "$(GRAFANA_PLUGIN_ID)"
	cd build && md5sum "$(GRAFANA_PLUGIN_ARTIFACT)" > "$(GRAFANA_PLUGIN_ARTIFACT_CHECKSUM)"
	rm -r "build/$(GRAFANA_PLUGIN_ID)"

##@ Helpers

clean: ## Clean all artifacts
	rm -rf .antlr .yarn build dist node_modules $(RMF_QUERY_PARSER_PATH)/*

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
