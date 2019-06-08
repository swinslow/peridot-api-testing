# SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

FROM golang:1.12

RUN mkdir -p /obsidian-api-testing
WORKDIR /obsidian-api-testing

ADD . /obsidian-api-testing

RUN go get -v ./...
RUN go build
