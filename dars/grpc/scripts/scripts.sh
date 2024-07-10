#!/bin/bash

proto_dir="protos"
output_base_dir="./genproto"

# Find all proto files in the specified directory
proto_files=($(find "$proto_dir" -name '*.proto'))

# Iterate over each proto file
for proto_file in "${proto_files[@]}"; do
  # Extract the base name of the proto file without extension
  base_name=$(basename "$proto_file" .proto)
  
  # Create the output directory for this proto file
  output_dir="$output_base_dir/$base_name"
  mkdir -p "$output_dir"
  
  # Run protoc command for this proto file
  protoc --go_out="$output_dir" --go_opt=paths=source_relative --go-grpc_out="$output_dir" --go-grpc_opt=paths=source_relative "$proto_file"
done
