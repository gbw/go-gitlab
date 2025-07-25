#!/usr/bin/env sh

case $(uname -s) in
Darwin)
  READLINK_FLAG="-f"
  ;;
*)
  READLINK_FLAG="-e"
  ;;
esac

testing_pkg_dir=$(readlink $READLINK_FLAG "$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)/../testing")
testing_client_generated_file="$testing_pkg_dir/client_generated.go"

instantiations_file="$(mktemp)"
client_fields_file="$(mktemp)"
testclient_mock_fields_file="$(mktemp)"
testclient_mocks_file="$(mktemp)"
# shellcheck disable=SC2162,SC2038
grep -E '^\s[A-Z][a-zA-Z0-9]+\s+[A-Z][a-zA-Z0-9]+Interface$' -- gitlab.go | while read line; do
  field=$(echo "$line" | awk '{ print $1 }')
  interface=$(echo "$line" | awk '{ print $2 }')

  echo "mock${field} := NewMock${interface}(ctrl)" >>"$instantiations_file"
  echo "${field}: mock${field}," >>"$client_fields_file"
  echo "Mock${field}: mock${field}," >>"$testclient_mock_fields_file"
  echo "Mock${field} *Mock${interface}" >>"$testclient_mocks_file"
done

(
  echo '// This file is generate from scripts/generate_testing_client.sh'
  echo 'package testing'
  echo ''
  echo 'import ('
  echo '   "fmt"'
  echo ''
  echo '   "go.uber.org/mock/gomock"'
  echo ''
  echo '   gitlab "gitlab.com/gitlab-org/api/client-go"'
  echo ')'
  echo ''
  echo 'type testClientMocks struct {'
  cat "$testclient_mocks_file"
  echo '}'
  echo ''
  echo 'func newTestClientWithCtrl(ctrl *gomock.Controller, options ...gitlab.ClientOptionFunc) *TestClient {'
  cat "$instantiations_file"
  echo ''
  echo '  c := &gitlab.Client{'
  cat "$client_fields_file"
  echo '  }'
  echo '  // Apply any given client options.'
  echo '  for _, fn := range options {'
  echo '    if fn == nil {'
  echo '      continue'
  echo '    }'
  echo '    if err := fn(c); err != nil {'
  echo '      panic(fmt.Errorf("a client option passed to the test client returned an error. This is not allowed. Error: %w", err))'
  echo '    }'
  echo '  }'
  echo ''
  echo '  return &TestClient{'
  echo '    Client: c,'
  echo '    testClientMocks: &testClientMocks{'
  cat "$testclient_mock_fields_file"
  echo '    },'
  echo '  }'
  echo '}'
) >"$testing_client_generated_file"

make fmt
