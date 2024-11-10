#!/usr/bin/env bash

# Applies the specified schema file to the database, similar to Terraform.
atlas schema apply --url "postgres://admin:admin@localhost:5433/happened_db?sslmode=disable" --to "file://sql/schema.sql" --dev-url "docker://postgres/15"
