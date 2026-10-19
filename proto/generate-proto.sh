#!/usr/bin/env bash

protoc \
  --proto_path=. \
  --go_out=. \
  --go_opt=module=github.com/git-sudo-404/Raft_Consensus \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/git-sudo-404/Raft_Consensus \
  proto/raft.proto
