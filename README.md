# bump

Bump is a CLI semantic versioning tool similar to `npm version`. It reads the latest git tag and increase the major, minor or patch part, and create a new tag to git.

## Install

```
go install github.com/neefrankie/bump
```

## Usage

```shell
bump major <-m message> # Increase major version
bump minor <-m message> # Increase minor version
bump patch <-m message> # Increase patch version
```

If `-m` flags is provided, an annotated tag will be created and the message will be passed to `git tag -m`; otherwise a lightweight tag will be created.

If the value of  `-m` flags contains `%s`, it will be replaced by the updated version.

