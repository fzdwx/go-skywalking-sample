#!/usr/bin/env just --justfile

build:
    go build -toolexec="/home/like/workspaces/helloworld/cmd/helloworld/agent -config /home/like/workspaces/helloworld/cmd/helloworld/config.yaml" -a  ./cmd/helloworld

build2:
    go build -toolexec="/home/like/workspaces/helloworld/cmd/helloworld/agent -config /home/like/workspaces/helloworld/cmd/helloworld/config.yaml" -a  ./cmd/qwe

update:
  go get -u
  go mod tidy -v
