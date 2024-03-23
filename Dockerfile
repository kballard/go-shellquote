# syntax=docker/dockerfile:1.2
# File : Dockerfile

ARG DEBIAN_RELEASE
ARG GOLANG_VERSION

# ------------------------------------------------------------------------------
FROM debian:${DEBIAN_RELEASE}-slim AS base-pure

# ------------------------------------------------------------------------------
FROM golang:${GOLANG_VERSION}-${DEBIAN_RELEASE} AS base-golang

# ------------------------------------------------------------------------------
FROM golangci/golangci-lint:v1.56 AS golangci-lint

# ------------------------------------------------------------------------------
FROM base-golang AS golang-init

WORKDIR /tmp/build
RUN apt-get update && apt-get install -y git
COPY go.mod go.sum ./
RUN go mod download

# ------------------------------------------------------------------------------
FROM golang-init AS golang-build

ARG BUILD_ENV
ARG BUILD_FLAGS
ARG BUILD_TARGET_NAME
ARG SRC
WORKDIR /tmp/build
COPY . .
RUN export ${BUILD_ENV}; go env; go build ${BUILD_FLAGS} -o ${BUILD_TARGET_NAME} ${SRC}

# ------------------------------------------------------------------------------
FROM golang-init AS golang-check

RUN go install gotest.tools/gotestsum@latest
COPY --from=golangci-lint /usr/bin/golangci-lint /usr/bin

# End of Dockerfile
