APP := bump

git_tag := `git tag -l --sort=-v:refname | head -n 1`
build_time := `date +%FT%T%z`
commit := `git log --max-count=1 --pretty=format:%aI_%h`
src_dir := .
build_dir := build

ldflags := -ldflags "-w -s -X github.com/neefrankie/bump/cmd.Version=$(git_tag)"

executable := $(build_dir)/$(APP)

run_generate := go run internal/version/main.go

.PHONY: build
build :
	@echo Building $(git_tag)
	go build -o $(executable) $(ldflags) -v $(src_dir)

.PHONY: run
run :
	./$(executable)

.PHONE:
major :
	$(run_generate) -major

.PHONY: minor
minor :
	$(run_generate) -minor

.PHONY: patch
patch :
	$(run_generate) -patch

.PHONY: clean
clean :
	go clean -x
	rm build/*
