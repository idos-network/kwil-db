#!/usr/bin/env bash

# How to run manually:
# docker build --pull --rm -f "contrib/devtools/Dockerfile" -t cosmossdk-proto:latest "contrib/devtools"
# docker run --rm -v $(pwd):/workspace --workdir /workspace cosmossdk-proto sh ./scripts/protocgen.sh

set -e

echo "Generating gogo proto code"
# cd proto
buf generate --template proto/buf.gen.gogo.yaml --path proto/kwil/node/v1/node.proto -o api/protobuf/
# proto_dirs=$(find ./kwil/node -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
# for dir in $proto_dirs; do
#   for file in $(find "${dir}" -maxdepth 2 -name '*.proto'); do
#     # this regex checks if a proto file has its go_package set to cosmossdk.io/api/...
#     # gogo proto files SHOULD ONLY be generated if this is false
#     # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
#     if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*cosmossdk.io/api' "$file" | grep -q ':0$'; then
#       buf generate --template buf.gen.gogo.yaml $file
#     fi
#   done
# done

# cd ..
mkdir -p api/protobuf/node/v1
mv api/github.com/kwilteam/kwil-db/api/protobuf/node/v1/* api/protobuf/node/v1/
rm -rf api/github.com
cp  api/protobuf/node/v1/* internal/node/
# # generate codec/testdata proto code
# (cd testutil/testdata; buf generate)

# # generate baseapp test messages
# (cd baseapp/testutil; buf generate)

# # move proto files to the right places
# cp -r github.com/cosmos/cosmos-sdk/* ./
# cp -r cosmossdk.io/** ./
# rm -rf github.com cosmossdk.io

# go mod tidy

# ./scripts/protocgen-pulsar.sh