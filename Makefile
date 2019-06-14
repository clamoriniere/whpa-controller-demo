PROJECT_NAME=whpa
ARTIFACT=controller
ARTIFACT_PLUGIN=kubectl-${PROJECT_NAME}

# 0.0 shouldn't clobber any released builds
TAG=v0.0.1
PREFIX =  ${PROJECT_NAME}/${ARTIFACT}
SOURCEDIR = "."

SOURCES := $(shell find $(SOURCEDIR) ! -name "*_test.go" -name '*.go')

all: build

vendor:
	go mod vendor

build: ${ARTIFACT}

${ARTIFACT}: ${SOURCES}
	CGO_ENABLED=0 go build -i -installsuffix cgo -ldflags '-w' -o ${ARTIFACT} ./cmd/manager/main.go

build-plugin: ${ARTIFACT_PLUGIN}

${ARTIFACT_PLUGIN}: ${SOURCES}
	CGO_ENABLED=0 go build -i -installsuffix cgo -ldflags '-w' -o ${ARTIFACT_PLUGIN} ./cmd/${ARTIFACT_PLUGIN}/main.go

container: vendor
	operator-sdk build $(PREFIX):$(TAG)
    ifeq ($(KINDPUSH), true)
	 kind load docker-image $(PREFIX):$(TAG)
    endif

test:
	./go.test.sh

e2e: container
	operator-sdk test local ./test/e2e --image $(PREFIX):$(TAG)

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f ${ARTIFACT}

validate:
	${GOPATH}/bin/golangci-lint run ./...

generate:
	operator-sdk generate k8s
	operator-sdk generate openapi

CRDS = $(wildcard deploy/crds/*_crd.yaml)
local-load: $(CRDS)
		for f in $^; do kubectl apply -f $$f; done
		kubectl apply -f deploy/
		kubectl delete pod -l name=${PROJECT_NAME}

$(filter %.yaml,$(files)): %.yaml: %yaml
	kubectl apply -f $@

install-tools:
	./hack/golangci-lint.sh -b ${GOPATH}/bin v1.16.0
	./hack/install-operator-sdk.sh

.PHONY: vendor build push clean test e2e validate local-load install-tools list