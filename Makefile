.PHONY: install release

ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=sachaos

install:
	go install

release:
	GOOS=windows GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/binder_windows_amd64
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/binder_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/binder_linux_amd64
	ghr -u $(GITHUB_USERNAME) -t $(shell cat github_token) --replace ${VERSION} $(ARTIFACTS_DIR)
