# Labeler [![Build Status](https://travis-ci.org/tonglil/labeler.svg?branch=master)](https://travis-ci.org/tonglil/labeler)

![logo](http://i.imgur.com/5wOQl2m.png)

Label management (create/rename/update/delete) on Github as code.

- [x] Using GitHub?
- [x] Want to commit/copy/share your label configuration?
- [ ] Use `labeler`!

For FOSS maintainers, enable your users to submit PRs and improve the process/label system!
- [Clean up][adobe] your labels.
- Move labels out of the [same][iconic] [flat][certbot] [space][ghost].
- Enforce a label color scheme that is not [meaningless][node] nor [confusing][babel] to view.

Inspired by [infrastructure as code][iac] tools like [Terraform][terraform] and organized label systems in projects like these:
- https://github.com/kubernetes/kubernetes/labels
- https://github.com/coreos/etcd/labels
- https://github.com/coreos/rkt/labels
- https://github.com/spf13/hugo/labels
- https://github.com/docker/docker/labels

[adobe]: https://github.com/adobe/brackets/labels
[iconic]: https://github.com/driftyco/ionic/labels
[certbot]: https://github.com/certbot/certbot/labels
[ghost]: https://github.com/TryGhost/Ghost/labels
[node]: https://github.com/nodejs/node/labels
[babel]: https://github.com/babel/babel/labels

[iac]: http://martinfowler.com/bliki/InfrastructureAsCode.html
[terraform]: https://github.com/hashicorp/terraform

## Installation

Get binaries for OS X / Linux / Windows from the latest [release].

Or use `go get`:

```
go get -u github.com/tonglil/labeler
```

[release]: https://github.com/tonglil/labeler/releases

## Usage

First, set a [GitHub token][tokens] in the environment (optional, the token can be set as an cli argument as well).

```
export GITHUB_TOKEN=xxx
```

> - The token for public repos need the `public_repo` scope.
> - The token for private repos need the `repo` scope.

[tokens]: https://github.com/settings/tokens

### Scanning labels

To scan existing labels from a repository and save it to a file:
```
labeler scan labels.yaml --repo owner/name
```

Which when run against a "new" repo created on GitHub, will:
- Fetch `bug` with color `fc2929`
- Fetch `duplicate` with color `cccccc`
- Fetch `enhancement` with color `84b6eb`
- Fetch `invalid` with color `e6e6e6`
- Fetch `question` with color `cc317c`
- Fetch `wontfix` with color `ffffff`

And write them into `labels.yaml`, creating the file if it doesn't exist, otherwise overwriting its contents.

### Applying labels

To apply labels to a repository:
```
labeler apply labels.yaml --dryrun
```

Where `labels.yaml` is like:
```yml
repo: owner/name
labels:
  - name: bug
    color: fc2929
  - name: help wanted
    color: 000000
  - name: fix
    color: cccccc
    from: wontfix
  - name: notes
    color: fbca04
```

Which when run against a "new" repo created on GitHub, will:
- Rename `wontfix` to `fix` with color `ffffff` to `ffffff`
- Update `help wanted` with color `159818` to `000000`
- Create `notes` with color `fbca04`
- Delete `duplicate` with color `cccccc`
- Delete `enhancement` with color `84b6eb`
- Delete `invalid` with color `e6e6e6`
- Delete `question` with color `cc317c`

When run again, rename changes will not be run because the label already exists.
In this manner, this tool is idempotent.

## Usage options

```
$ labeler
Labeler is a CLI application for managing labels on Github as code.

With the ability to scan and apply label changes, repository maintainers can
empower contributors to submit PRs and improve the project management
process/label system!

Usage:
  labeler [command]

Available Commands:
  apply       Apply a YAML label definition file
  completion  Output shell completion code for tab completion
  scan        Save a repository's labels into a YAML definition file
  version     Print the version information

Use "labeler [command] --help" for more information about a command.
```

## Tab completion

```bash
source <(labeler completion)
```

## Development

[`glide`][glide] is used to manage vendor dependencies.

Roadmap:
- Plan -> execute (aka always dry-run first).
- Automatically update file after renaming operations are complete.
- Organizational support (apply/only-add one config to all repos in an organization).

[glide]: https://github.com/Masterminds/glide

## Testing

**This could use your contribution!**
Help me create a runnable test suite.

## See also

- Rust: https://github.com/jimmycuadra/ghlabel
- Node: https://github.com/popomore/github-labels
- Node: https://github.com/repo-utils/org-labels
- PHP: https://gist.github.com/zot24/0cbbd3ee4b22123cb62a
