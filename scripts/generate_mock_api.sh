#!/usr/bin/env sh

# This script generates the `api_generated.go` file, which includes one `go:generate` line for each interface.
# This is used to ensure we have a mocked setup for each interface we create in the client.

case $(uname -s) in
Darwin)
  READLINK_FLAG="-f"
  ;;
*)
  READLINK_FLAG="-e"
  ;;
esac

testing_pkg_dir=$(readlink $READLINK_FLAG "$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)/../testing")
api_file="$testing_pkg_dir/api_generated.go"

(
  echo '// This file is generate from scripts/generate_mock_api.sh'
  echo 'package testing'
  echo ''
) >"$api_file"

(
# shellcheck disable=SC2162,SC2038
grep -E '^\s[A-Z][a-zA-Z0-9]+Interface interface {$' -- *.go | awk '{ print $1 $2 }' | while read line; do
  filename=$(echo "$line" | cut -d: -f1)
  filename=${filename%.go}
  interface=$(echo "$line" | cut -d: -f2)

  echo "//go:generate go run go.uber.org/mock/mockgen@v0.5.2 -typed -destination=${filename}_mock.go -package=testing gitlab.com/gitlab-org/api/client-go ${interface}"
done
) | LC_ALL=C sort >> "$api_file"

go generate ./...
make fmt
