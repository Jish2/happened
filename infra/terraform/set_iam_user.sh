#!/usr/bin/env bash

echo "executing aws profile script"

PROFILE=$(terraform output -raw iam_admin_username)
ACCESS_KEY_ID=$(terraform output -raw iam_admin_id)
SECRET_KEY=$(terraform output -raw iam_admin_secret)
REGION=$(terraform output -raw aws_region)

aws configure set aws_access_key_id "$ACCESS_KEY_ID" --profile "$PROFILE"
aws configure set aws_secret_access_key "$SECRET_KEY" --profile "$PROFILE"
aws configure set region "$REGION" --profile "$PROFILE"