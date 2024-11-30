#!/bin/bash

set -e

DEPLOYMENT_NAME="${DEPLOYMENT_NAME:-$1}"

if [ -z "$DEPLOYMENT_NAME" ]; then
  echo "ERROR: DEPLOYMENT_NAME is required."
  exit 1
fi
# Sanitize DEPLOYMENT_NAME
DEPLOYMENT_NAME=$(echo "$DEPLOYMENT_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')


gcloud run services delete "$DEPLOYMENT_NAME" --region us-west1 --quiet