# Versioning

A library to make set/get/writing versioning information easier.

## Usage

1. Add in `main.go`:

    ```go
    var version string

    func init() {
        versioning.Set(version)
    }
    ```

1. Where you want to get the version:

    ```go
    version := versioning.Get()
    // or
    versioning.Write(os.Stdout)
    ```

1. Build your binary with this flag:

    ```bash
    $ go build -ldflags "-X main.version=123-up-to-you"
    ```

## Development

Roadmap:

- Checking new versions
  - Possibly use semver: https://github.com/hashicorp/go-version or https://github.com/blang/semver
- Self-upgrading binary
