#!/bin/bash
set -e

DEPLOYMENT_NAME="${DEPLOYMENT_NAME:-$1}"

if [ -z "$DEPLOYMENT_NAME" ]; then
  echo "ERROR: DEPLOYMENT_NAME is required."
  exit 1
fi

# Clean the name to only have alphanumeric and slashes
DEPLOYMENT_NAME=$(echo "$DEPLOYMENT_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')

gcloud run services replace service.yaml --quiet
gcloud run services set-iam-policy "$DEPLOYMENT_NAME" policy.yaml --region us-west1 --quiet