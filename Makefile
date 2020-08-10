.PHONY: build

build:
	goreleaser release --snapshot --skip-publish --rm-dist

release-test:
	goreleaser release --skip-publish --rm-dist

release:
	goreleaser release --rm-dist