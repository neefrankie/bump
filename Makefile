APP := bump

version := `git tag -l --sort=-v:refname | head -n 1`
build_time := `date +%FT%T%z`
commit := `git log --max-count=1 --pretty=format:%aI_%h`
src_dir := .
build_dir := build

ldflags := -ldflags "-w -s -X github.com/neefrankie/bump/cmd.Version=$(version)"

executable := $(build_dir)/$(APP)

.PHONY: build
build :
	@echo Building $(version)
	go build -o $(executable) $(ldflags) -v $(src_dir)

.PHONY: run
run :
	./$(executable)

.PHONY: clean
clean :
	go clean -x
	rm build/*
