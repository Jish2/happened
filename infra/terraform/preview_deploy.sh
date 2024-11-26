#!/bin/bash

# deploy-cloud-run.sh
SERVICE_NAME=$1
IMAGE_URL=$2
PORT=${3:-8080}  # Default to 8080 if not provided
REGION=${4:-us-west1}  # Default to us-central1 if not provided

echo "Deploying ${SERVICE_NAME}..."

gcloud run deploy ${SERVICE_NAME} \
  --image=${IMAGE_URL} \
  --port=${PORT} \
  --region=${REGION} \
  --platform=managed


echo "Making service public..."
gcloud run services add-iam-policy-binding ${SERVICE_NAME} \
  --region=${REGION} \
  --member="allUsers" \
  --role="roles/run.invoker"