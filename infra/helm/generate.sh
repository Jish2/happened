#!/bin/bash
set -e

# Ensure all required environment variables are set
REQUIRED_ENV_VARS=("CLERK_SECRET_KEY" "DB_HOST" "DB_USER" "DB_PASS" "DB_NAME" "DB_PORT")
for var in "${REQUIRED_ENV_VARS[@]}"; do
  if [ -z "${!var}" ]; then
    echo "Error: Environment variable $var is not set."
    exit 1
  fi
done

# Ensure deployment name is provided as an argument
if [ -z "$1" ]; then
  echo "Error: Deployment name is required."
  exit 1
fi

DEPLOYMENT_NAME=$1
echo Deploying as "$DEPLOYMENT_NAME"


# Write secrets to values.yaml
VALUES_FILE="./cloudrun/values.yaml"
echo "DB_HOST: $DB_HOST" >> $VALUES_FILE
echo "DB_USER: $DB_USER" >> $VALUES_FILE
echo "DB_PASS: $DB_PASS" >> $VALUES_FILE
echo "DB_NAME: $DB_NAME" >> $VALUES_FILE
echo "DB_PORT: $DB_PORT" >> $VALUES_FILE 
echo "CLERK_SECRET_KEY: $CLERK_SECRET_KEY" >> $VALUES_FILE

echo "Updated $VALUES_FILE:"
cat "$VALUES_FILE"


COMMIT_HASH=$(git rev-parse --short HEAD)
# Sanitize DEPLOYMENT_NAME
DEPLOYMENT_NAME=$(echo "$DEPLOYMENT_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')

IMAGE="docker.io/anmho/happened:$COMMIT_HASH"
# Run Helm template command
helm template cloudrun \
    --debug \
    --set name="$DEPLOYMENT_NAME" \
    --set image="$IMAGE" > service.yaml