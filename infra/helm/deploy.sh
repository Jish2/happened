#!/bin/bash
set -e

BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
# Sanitize BRANCH_NAME
BRANCH_NAME=$(echo "$BRANCH_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')

SERVICE="happened-$BRANCH_NAME"
gcloud run services replace service.yaml
gcloud run services set-iam-policy $SERVICE policy.yaml