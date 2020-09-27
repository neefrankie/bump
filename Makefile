APP := bump

version := `git tag -l --sort=-v:refname | head -n 1`
build_time := `date +%FT%T%z`
commit := `git log --max-count=1 --pretty=format:%aI_%h`
src_dir := .
build_dir := build

ldflags := -ldflags "-w -s -X github.com/neefrankie/bump/cmd.Version=$(version)"

is_git_clean = `git diff --stat`

executable := $(build_dir)/$(APP)

run_generate := go run internal/version/generate.go

.PHONY: build
build :
	@echo Building $(version)
	go build -o $(executable) $(ldflags) -v $(src_dir)

.PHONY: run
run :
	./$(executable)

.PHONY: publish-major
publish-major : build
ifeq ($(strip $(is_git_clean)),)
	$(run_generate) -major
	git add . && git commit -m "Major version"
	$(executable) major
	git push && git push --tags
else
	@echo WARNING: Repository is not clean. Please commit untracked files.
endif

.PHONY: publish-minor
publish-minor :
ifeq ($(strip $(is_git_clean)),)
	$(run_generate) -minor
	git add . && git commit -m "Minor version"
	$(executable) minor
	git push && git push --tags
else
	@echo WARNING: Repository is not clean. Please commit untracked files.
endif

.PHONY: publish-patch
publish-patch :
ifeq ($(is_git_clean),)
#	$(run_generate) -patch
#	git add . && git commit -m "Patch"
#	$(executable) patch
#	git push && git push --tags
	@echo Clean
else
	@echo WARNING: Repository is not clean. Please commit untracked files.
endif

.PHONY: clean
clean :
	go clean -x
	rm build/*
