#!/bin/bash
set -euo pipefail
set -x

main() {
  if [ "${TRAVIS_BRANCH}" == "master" ] || [[ "${TRAVIS_BRANCH}" =~ ${LABELS_BRANCH:-^labels-ci} ]]; then
    go get -u github.com/tonglil/labeler
    labeler apply labels.yaml -l 9

    if [ "${LABELS_TRIM:-0}" -eq 1 ]; then
      labeler trim labels.yaml

      diff=$(git diff --name-only -- labels.yaml)
      if [ "$diff" == "labels.yaml" ]; then
        config_ssh
        config_git

        branch="ci-trim-labels_${TRAVIS_BRANCH}_$(date +%Y-%m-%d-%H-%M-%S)"
        git checkout -b "$branch"
        git add labels.yaml
        git commit -m 'GENERATED BY CI: trim labels.yaml'
        git push --set-upstream ci "$branch"

        if [ "${LABELS_PR:-0}" -eq 1 ]; then
          config_pr
          go get -u github.com/github/hub
          hub pull-request -F PATCH.md -b "${TRAVIS_BRANCH}" -h "$branch"
        fi
      fi
    fi
  fi
}

# https://docs.travis-ci.com/user/deployment/custom/
# Start the ssh agent
config_ssh() {
  eval "$(ssh-agent -s)"
  # Decrypted in .travis.yml, this private deploy key should have push access
  chmod 600 deploy_key
  ssh-add deploy_key
}

config_git() {
  git config user.email "travis@ci"
  git config user.name "Travis CI"

  git remote add ci "git@github.com:$TRAVIS_REPO_SLUG.git"
}

config_pr() {
  cat <<EOF > PATCH.md
[ci] Trim labels.yaml

Automatically generated by \`labeler trim labels.yaml\`.
https://github.com/tonglil/labeler.
EOF
}

main