#!/usr/bin/env sh

# This script generates a map keyed from the Service structs to the interfaces they implement. This is used
# to test that each of the service structs implement their interface properly, and automatically keeps the
# map up to date so that contributors who implement new services don't need to do as much manual work.
case $(uname -s) in
Darwin)
  READLINK_FLAG="-f"
  ;;
*)
  READLINK_FLAG="-e"
  ;;
esac

root_dir=$(readlink $READLINK_FLAG "$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)/..")
api_service_map_test_file="$root_dir/gitlab_service_map_generated_test.go"

(
  echo '// This file is generate from scripts/generate_service_interface_map.sh'
  echo 'package gitlab'
  echo ''
  echo 'var ('
  echo '  serviceMap = map[any]any{'
) >"$api_service_map_test_file"

(
# shellcheck disable=SC2162,SC2038
grep -E '^\s[A-Z][a-zA-Z0-9]+Service struct {' -- *.go | awk '{ print $1 $2 }' | while read line; do
  filename=$(echo "$line" | cut -d: -f1)
  filename=${filename%.go}
  service=$(echo "$line" | cut -d: -f2)

  echo "&${service}{}: (*${service}Interface)(nil),"
done
) | sort >> "$api_service_map_test_file"

(
  echo '  }'
  echo ')'
) >>"$api_service_map_test_file"

make fmt
