#!/usr/bin/env bash

if ! id -u "$UID" >/dev/null 2>/dev/null; then
    adduser -D -u "$UID" myuser 2>/dev/null
fi

if ! getent group "$GID" >/dev/null 2>&1; then
    addgroup -g "$GID" mygroup
fi
#
#echo "Resolving references"
#
#redocly bundle -d /opt/in/spec.yaml -o /opt/in/dereferenced.yaml

echo "Generating files"

oapi-codegen -package oapi -generate spec /opt/in/spec.yaml > /opt/out/openapi.go
oapi-codegen -package oapi -generate client /opt/in/spec.yaml > /opt/out/client.go
oapi-codegen -package oapi -generate types /opt/in/spec.yaml > /opt/out/types.go

echo "Setting owner user/group"

chown "$UID":"$GID" -R /opt/out/*

echo "Done"
