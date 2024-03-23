# File : Makefile

PROJECT_NAME := go-shellquote
PWD := $(shell pwd)

DEBIAN_RELEASE ?= bookworm
GOLANG_VERSION ?= 1.21

export DOCKER_BUILDKIT=1

LINT := golangci-lint run --max-issues-per-linter 0 --max-same-issues 0 --verbose ./...
LINT_FIX := golangci-lint run --verbose --fix
TEST_UNIT := gotestsum --format testname ./... -- -p 1

.PHONY: lint
## Run non destructive lint over code base.
lint: build-check
	$(info Run non destructive lint over codebase $(PWD))
	docker run --volume $(PWD):/tmp/build/:ro $(PROJECT_NAME):check $(LINT)

.PHONY: test
## Run all available test suites.
test: test-unit test-integration test-smoke
	$(info Run all available test suites)

.PHONY: test-unit
## Run available unittests, access to db/es instances is required
test-unit: build-check
	$(info Run unit available unittests)
	docker run --volume $(PWD):/tmp/build/:ro $(PROJECT_NAME):check $(TEST_UNIT)

.PHONY: test-integration
## Run integration test with external services
test-integration:
	$(info Run integration test with external services)
	echo TBA

.PHONY: test-smoke
## Run smoke test with external services
test-smoke:
	$(info Run smoke test with external services)
	echo TBA

DOCKER_BUILD_ARGS ?= \
	--build-arg DEBIAN_RELEASE=$(DEBIAN_RELEASE) \
	--build-arg GOLANG_VERSION=$(GOLANG_VERSION)

.PHONY: test-init
## Build isolated test container to use in all verity of tests.
build-check:
	$(info Build isolated test container to use in all verity of tests)
	docker build \
	$(DOCKER_BUILD_ARGS) \
	--file Dockerfile \
	--tag $(PROJECT_NAME):check \
	--target=golang-check \
	.

# End of Makefile
