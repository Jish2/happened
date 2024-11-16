#!/usr/bin/env bash

# Generate the SDK asynchronously

./bin/api openapi > openapi.yaml
make -C ../ gen

./bin/api
