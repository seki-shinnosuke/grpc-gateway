#!/bin/bash

readonly protoc_out_dir='gen'
readonly protoc_src_dir='proto'
readonly protoc_include_dir="${protoc_src_dir}/include"

readonly go_out_dir="${protoc_out_dir}/go"
readonly openapiv2_out_dir="${protoc_out_dir}/openapiv2"

cleanup_protoc_output_dirs() {
  rm -rf "${go_out_dir}" "${openapiv2_out_dir}"
  mkdir -p "${go_out_dir}" "${openapiv2_out_dir}"
}

execute_protoc() {
  local go_prefix="github.com/seki-shinnosuke/grpc-gateway/gen/go"

  find "${protoc_src_dir}" -name '*.proto' -print0 \
    | xargs -0 protoc -I "${protoc_src_dir}" -I "${protoc_include_dir}" \
    --go_out="${go_out_dir}" \
    --go_opt=module="${go_prefix}" \
    --go-grpc_out="${go_out_dir}" \
    --go-grpc_opt=module="${go_prefix}" \
    --grpc-gateway_out="${go_out_dir}" \
    --grpc-gateway_opt=module="${go_prefix}" \
    --openapiv2_out="${openapiv2_out_dir}" \
    --openapiv2_opt=allow_merge=true
}

cleanup_protoc_output_dirs
execute_protoc
go mod tidy