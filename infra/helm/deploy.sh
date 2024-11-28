#!/bin/bash
set -e

BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
# Sanitize BRANCH_NAME
BRANCH_NAME=$(echo "$BRANCH_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')
echo "BRANCH_NAME=$BRANCH_NAME"

SERVICE="happened-$BRANCH_NAME"
echo "SERVICE=$SERVICE"
gcloud run services replace service.yaml --quiet
gcloud run services set-iam-policy "$SERVICE" policy.yaml --region us-west1 --quiet