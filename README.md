# gh-action-escape

[![CI - Go](https://github.com/kachick/gh-action-escape/actions/workflows/ci-go.yml/badge.svg?branch=main)](https://github.com/kachick/gh-action-escape/actions/workflows/ci-go.yml?query=event%3Apush++)
[![CI - E2E](https://github.com/kachick/gh-action-escape/actions/workflows/ci-e2e.yml/badge.svg)](https://github.com/kachick/gh-action-escape/actions/workflows/ci-e2e.yml)
[![CI - Nix](https://github.com/kachick/gh-action-escape/actions/workflows/ci-nix.yml/badge.svg?branch=main)](https://github.com/kachick/gh-action-escape/actions/workflows/ci-nix.yml?query=branch%3Amain+)

Escape/Wrap given multiline text with random delimiter for `$GITHUB_OUTPUT` and `$GITHUB_ENV`

See [official docs](https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#example-of-a-multiline-string) for the background

## Usage

Before

```yaml
steps:
  # You should write these handlings in all steps that use multiline value with GITHUB_OUTPUT and/or GITHUB_ENV
  - name: Set the value in bash
    id: step_one
    run: |
      EOF=$(dd if=/dev/urandom bs=15 count=1 status=none | base64)
      echo "json<<$EOF" >> "$GITHUB_OUTPUT"
      curl https://example.com >> "$GITHUB_OUTPUT"
      echo "$EOF" >> "$GITHUB_OUTPUT"
  - name: Use product in a before step
    run: echo "The result is ${{ steps.step_one.outputs.json }}"
```

After

```yaml
steps:
  # Once installed, the cli can be used in all following steps
  - name: Install gh-action-escape
    run: curl -fsSL https://raw.githubusercontent.com/kachick/gh-action-escape/main/scripts/install-in-github-action.sh | sh -s v0.2.0
  - name: Set the value in bash
    id: step_one
    run: curl https://example.com | gh-action-escape -name=json >> "$GITHUB_OUTPUT"
  - name: Use product in a before step
    run: echo "The result is ${{ steps.step_one.outputs.json }}"
```

## Installation

[Prebuilt binaries](https://github.com/kachick/gh-action-escape/releases)

```console
> curl -L https://github.com/kachick/gh-action-escape/releases/latest/download/gh-action-escape_Linux_x86_64.tar.gz | tar xvz -C ./ gh-action-escape
> ./gh-action-escape --version
...
```

In [Nix](https://nixos.org/), you can skip installation steps

```console
> nix run github:kachick/gh-action-escape -- --version
gh-action-escape dev (rev) # unknown
> nix run github:kachick/gh-action-escape/v0.2.0 -- --help
...(With specific version)
```

`go install`

```console
> go install github.com/kachick/gh-action-escape/cmd/gh-action-escape@latest
go: downloading...
> ${GOPATH:-"$HOME/go"}/bin/gh-action-escape --version
gh-action-escape dev (rev) # unknown
```

## Motivation

Just for <https://github.com/kachick/depop/pull/27/files>
