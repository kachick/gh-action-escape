#!/bin/sh

set -e

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

mkdir -p /home/runner/.gh-action-multiline/bin
curl -L "https://github.com/kachick/gh-action-multiline/releases/download/v0.1.0/gh-action-multiline_${suffix}" | tar xvz -C /home/runner/.gh-action-multiline/bin gh-action-multiline
chmod +x /home/runner/.gh-action-multiline/bin/gh-action-multiline
echo '/home/runner/.gh-action-multiline/bin' >> "$GITHUB_PATH"
