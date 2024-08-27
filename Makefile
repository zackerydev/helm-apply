HELM_HOME ?= $(shell helm home)
VERSION := $(shell sed -n -e 's/version:[ "]*\([^"]*\).*/\1/p' plugin.yaml)

HELM_3_PLUGINS := $(shell helm env HELM_PLUGINS)

PKG := github.com/zackerydev/helm-apply
LDFLAGS := -X $(PKG)/cmd.Version=$(VERSION)

GO ?= go


.PHONY: install
install: build
	mkdir -p $(HELM_3_PLUGINS)/helm-apply/bin
	cp bin/apply $(HELM_3_PLUGINS)/helm-apply/bin
	cp plugin.yaml $(HELM_3_PLUGINS)/helm-apply/

.PHONY: build
build:
	mkdir -p bin/
	go build -v -o bin/apply -ldflags="$(LDFLAGS)"
