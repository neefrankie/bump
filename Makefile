APP := bump

git_tag := `git tag -l --sort=-v:refname | head -n 1`
build_time := `date +%FT%T%z`
commit := `git log --max-count=1 --pretty=format:%aI_%h`
src_dir := .
build_dir := build

ldflags := -ldflags "-w -s -X github.com/neefrankie/bump/cmd.Version=$(git_tag)"

executable := $(build_dir)/$(APP)

run_generate := go run internal/version/generate.go

.PHONY: build
build :
	@echo Building $(git_tag)
	go build -o $(executable) $(ldflags) -v $(src_dir)

.PHONY: run
run :
	./$(executable)

.PHONE:
version-major :
	$(run_generate) -major

.PHONY: publish-major
publish-major : version-major build
	git add . && git commit -m "Bump version `cat build/version.txt`"
	$(executable) major
	git push && git push --tags

.PHONY: version-minor
version-minor :
	$(run_generate) -minor

.PHONY: publish-minor
publish-minor : version-minor build
	git add . && git commit -m "Bump version `cat build/version.txt`"
	$(executable) minor
	git push && git push --tags

.PHONY: version-patch
version-patch :
	$(run_generate) -patch

.PHONY: publish-patch
publish-patch : version-patch build
	git add . && git commit -m "Bump version `cat build/version.txt`"
	$(executable) patch
	git push && git push --tags

.PHONY: clean
clean :
	go clean -x
	rm build/*
