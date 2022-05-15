GO        ?= go
TAGS      :=
LDFLAGS	  := -w -s
TESTS     := .
TESTFLAGS :=
GOFLAGS   :=
GOMOD     :=
RUNFLAGS  :=
BINDIR    := $(CURDIR)/bin

all: commander

.PHONY: setup
setup:
	go get -u sigs.k8s.io/controller-tools/cmd/controller-gen

.PHONY: run-in-minikube
run-in-minikube:
	eval $(minikube -p minikube docker-env) && \
	docker build -t testitivity . && \
	kubectl run testitivity --rm -i --tty --image testitivity

commander: setup
	go generate ./...
	CGO_ENABLED=0 GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]') GOARCH=amd64 $(GO) build -a -gcflags='all=-N -l' -ldflags="$(LDFLAGS)" -v $(GOMOD)./cmd/commander/commander.go;
