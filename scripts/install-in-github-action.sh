#!/bin/bash

set -euo pipefail

case $(uname -sm) in
"Linux x86_64") suffix="Linux_x86_64.tar.gz" ;;
"Linux aarch64") suffix="Linux_arm64.tar.gz" ;;
"Linux i386|Linux i586|Linux i686") suffix="Linux_i386.tar.gz" ;;
"Darwin x86_64") suffix="Darwin_x86_64.tar.gz" ;;
"Darwin arm64") suffix="Darwin_arm64.tar.gz" ;;
*)
	echo "does not support this machine: $(uname -sm)"
	exit 1
esac

if [ $# -eq 0 ]; then
	archive_uri="https://github.com/kachick/gh-action-escape/releases/latest/download/gh-action-escape_${suffix}"
else
	archive_uri="https://github.com/kachick/gh-action-escape/releases/download/${1}/gh-action-escape_${suffix}"
fi

install_in="${XDG_DATA_HOME:-$HOME}/.gh-action-escape/bin"

mkdir -p "$install_in"
curl -L "$archive_uri" | tar xvz -C "$install_in" gh-action-escape
chmod +x "${install_in}/gh-action-escape"
echo "$install_in" | tee -a "$GITHUB_PATH"
