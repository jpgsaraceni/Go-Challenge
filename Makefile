.PHONY: metalint
metalint:
ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
endif
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners -c ./.golangci.yml ./...