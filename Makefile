include Makefile.vars

.PHONY: default
default: test

.PHONY: install.gvm
install.gvm:
	@echo "$(REPO) install.gvm"
	which gvm || \
		curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash

.PHONY: install
install: deps vendor
	@echo "$(REPO) install"

.PHONY: deps
deps:
	@echo "$(REPO) deps"
	@which gotestsum || \
		curl -O -L https://github.com/gotestyourself/gotestsum/releases/download/v0.3.2/gotestsum_0.3.2_linux_amd64.tar.gz && \
		tar xf gotestsum_0.3.2_linux_amd64.tar.gz && \
		mv gotestsum /usr/local/bin
	gotestsum --help > /dev/null 2>&1
	which dlv || \
		go get -u github.com/derekparker/delve/cmd/dlv
	dlv version

.PHONY: vendor
vendor:
	@echo "$(REPO) vendor"
	GOCACHE=on GO111MODULE=on go mod vendor
	GOCACHE=on GO111MODULE=on go mod verify

.PHONY: clean
clean:
	@echo "$(REPO) clean"
	-rm $(NAME)*coverage*
	-rm *.test
	-rm *.pprof

.PHONY: local
local:
	@echo "Set enviroment to local"

.PHONY: dev
dev:
	@echo "Set enviroment to dev"

.PHONY: staging
staging:
	@echo "Set enviroment to staging"

.PHONY: prod
prod:
	@echo "Set enviroment to prod"

.PHONY: checkenv
checkenv:
	@echo "$(REPO) checkenv"
ifeq ($(ENV), )
	echo "err_blank_env: env=$(ENV)"
	exit 540
endif

.PHONY: all
all: clean vet coverage.text bench

.PHONY: fmt
fmt:
	@echo "$(REPO) fmt"
	go fmt $(PKGS)

.PHONY: vet
vet:
	@echo "$(REPO) vet"
	go vet $(PKGS)

.PHONY: debug
debug:
	@echo "$(REPO) debug"
	dlv debug $(REPO)

.PHONY: debugtest
debugtest:
	@echo "$(REPO) debugtest"
ifeq ($(TEST_PKGS),)
	@echo "debugtest: pkgs=*"
	dlv test --build-flags='$(REPO)' -- -test.run $(TESTS)
else
	@echo "debugtest: pkgs=$(TEST_PKGS)"
	dlv test --build-flags='$(REPO)/$(TEST_PKGS)' -- -test.run $(TESTS)
endif

.PHONY: test
test:
	@echo "$(REPO) test"
ifeq ($(TEST_PKGS),)
	@echo "test: pkgs=*"
	gotestsum -f short-verbose -- -v -race -run $(TESTS) $(PKGS)
else
	@echo "test: pkgs=$(TEST_PKGS)"
	$(foreach pkg,$(TEST_PKGS),\
		gotestsum -f short-verbose -- -v -race -run $(TESTS) $(REPO)/$(pkg);\
	)
endif

.PHONY: itest
itest:
	@echo "$(REPO) itest"
ifeq ($(TEST_PKGS),)
	@echo "itest: pkgs=*"
	gotestsum -f short-verbose -- -tags=integration -v -race -run $(TESTS) $(PKGS)
else
	@echo "itest: pkgs=$(TEST_PKGS)"
	$(foreach pkg,$(TEST_PKGS),\
		gotestsum -f short-verbose -- -tags=integration -v -race -run $(TESTS) $(REPO)/$(pkg);\
	)
endif

.PHONY: bench
bench:
	@echo "$(REPO) bench"
ifeq ($(TEST_PKGS),)
	@echo "bench: pkgs=*"
	gotestsum -f short-verbose -- -bench=. -run="^$$" -benchmem $(PKGS)
else
	@echo "bench: pkgs=$(TEST_PKGS)"
	$(foreach pkg,$(TEST_PKGS),\
		gotestsum -f short-verbose -- -bench=. -run="^$$" -cpuprofile=cpu.pprof -memprofile=mem.pprof -benchmem $(REPO)/$(pkg);\
	)
endif

.PHONY: coverage
coverage:
	@echo "$(REPO) coverage"
	@touch $(COVERAGE_FILE)
ifeq ($(TEST_PKGS),)
	@echo "coverage: pkgs=*"
	gotestsum -f short-verbose -- -tags=integration -v -run $(TESTS) -coverpkg=./... -coverprofile=$(COVERAGE_FILE) $(PKGS)
else
	@echo "bench: pkgs=$(TEST_PKGS)"
	@touch $(COVERAGE_PKG_FILE)
	@echo 'mode: set' > $(COVERAGE_FILE)
	$(foreach pkg,$(TEST_PKGS),\
		gotestsum -f short-verbose -- -tags=integration -v -run $(TESTS) -coverpkg=./... -coverprofile=$(COVERAGE_PKG_FILE) $(REPO)/$(pkg);\
		grep -v 'mode: set' $(COVERAGE_PKG_FILE) >> $(COVERAGE_FILE);\
	)
endif

.PHONY: coverage.text
coverage.text: coverage
	@echo "$(REPO) coverage.text"
	go tool cover -func=$(COVERAGE_FILE)

.PHONY: coverage.html
coverage.html: coverage
	@echo "$(REPO) coverage.html"
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	open $(COVERAGE_HTML) || google-chrome $(COVERAGE_HTML) || google-chrome-stable $(COVERAGE_HTML)

.PHONY: release
release: checkenv coverage.text docker
	@echo "$(REPO) release"

.PHONY: docker
docker:
	@echo "$(REPO)@$(BUILD) docker"
	docker build -t $(DOCKER_NAME) -t $(DOCKER_NAME):$(VERSION) -f ./etc/docker/Dockerfile .

.PHONY: docker.bash
docker.bash:
	@echo "$(REPO)@$(BUILD) docker.bash"
	docker run --rm --name $(NAME) --entrypoint bash -it -v `pwd`:/go/src/$(REPO) $(DOCKER_NAME)

docker.%:
	@echo "$(REPO)@$(BUILD) docker.$*"
	docker run --rm --name $(NAME) -v `pwd`:/go/src/$(REPO) $(DOCKER_NAME) $*
