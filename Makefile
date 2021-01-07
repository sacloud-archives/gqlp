#
# Copyright 2021 The gqlp Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
AUTHOR          ?="The gqlp Authors"
COPYRIGHT_YEAR  ?="2021"
COPYRIGHT_FILES ?=$$(find . -name "*.go" -print | grep -v "/vendor/")

default: gen fmt set-license goimports lint test

.PHONY: tools
tools:
	go get github.com/99designs/gqlgen
	go get golang.org/x/tools/cmd/goimports
	go get github.com/sacloud/addlicense
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.23.8/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.23.8


.PHONY: run
run:
	go run github.com/sacloud/gqlp/cmd/gqlp

.PHONY: test
test:
	go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 -race;

.PHONY: gen
gen: _gen fmt goimports set-license

.PHONY: _gen
_gen:
	go generate ./...

.PHONY: goimports
goimports: fmt
	goimports -l -w .

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: set-license
set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)

