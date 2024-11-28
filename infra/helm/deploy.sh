#!/bin/bash
set -e

BRANCH_NAME=$(git branch --show-current)
# Sanitize BRANCH_NAME
BRANCH_NAME=$(echo "$BRANCH_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')
echo "BRANCH_NAME=$BRANCH_NAME"

NAME="happened-$BRANCH_NAME"
echo "NAME=$NAME"
gcloud run services replace service.yaml --quiet
gcloud run services set-iam-policy "$NAME" policy.yaml --region us-west1 --quiet