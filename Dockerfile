# SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

FROM golang:1.12

RUN mkdir -p /peridot-api-testing
WORKDIR /peridot-api-testing

ADD . /peridot-api-testing

RUN go get -v ./...
RUN go build
