# How to develop

## Setup

1. Install [Nix](https://nixos.org/) package manager
2. Run `nix-shell` or `nix-shell --command 'zsh'`
3. You can use development tools

```console
> nix-shell
(prepared bash)

> task fmt
task: [fmt] dprint fmt
task: [fmt] go fmt ./...

> task
task: [build] ..."
task: [test] go test ./...
task: [lint] dprint check
task: [lint] go vet ./...
PASS
ok      gh-action-escape    0.313s

> find dist
dist
dist/metadata.json
dist/config.yaml
dist/gh-action-escape_linux_amd64_v1
dist/gh-action-escape_linux_amd64_v1/gh-action-escape
dist/artifacts.json

> ./dist/gh-action-escape_linux_amd64_v1/gh-action-escape --version
gh-action-escape 0.2.0-next (906924b) # 2023-06-19T09:33:14Z
```
