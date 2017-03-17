# Travis CI Automation

Automate `labels.yaml` application to the repository so they are applied when changes get merged.

## Usage

- [.travis.yml](.travis.yml): CI configuration file
- [apply.sh](apply.sh): automation script

Add the example configuration to your .travis.yml.

### 1. Environment variables

Set some environment variables:

- `secure: "xxx"`: `GITHUB_TOKEN` encrypted output from [`travis encrypt GITHUB_TOKEN="token"`][encryption]
- `LABELS_BRANCH`: a regex for the branches to perform labels automation (default `^labels-ci`)
- `LABELS_TRIM`: set to `1` to trim `labels.yaml`, commit changes into a new branch, and push to GitHub (default `0`)
- `LABELS_PR`: set to `1` to open a PR (default `0`)

[encryption]: https://docs.travis-ci.com/user/encryption-keys/

### 2. `openssl` command

To commit from Travis CI to GitHub when `LABELS_TRIM` or `LABELS_PR` is set, the repo needs a deploy key.

Generate a new public + private deploy key pair for this repo:

```
ssh-keygen -t rsa -b 4096 -C "travis@ci" -f deploy_key -N ''
```

Add the public key to GitHub with write access (`https://github.com/<username>/<repository>/settings/keys`).

Use [`travis encrypt-file deploy_key`][encrypt-file] and add the resulting `openssl` command to .travis.yml.

Commit `deploy_key.enc`.

Remove `deploy_key` and `deploy_key.pub`.

[encrypt-file]: https://docs.travis-ci.com/user/encrypting-files/

### 3. `apply.sh` command

Add this script to your repository.
