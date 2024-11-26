
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
# Sanitize BRANCH_NAME
BRANCH_NAME=$(echo "$BRANCH_NAME" | sed 's/[^a-zA-Z0-9-]/-/g')

NAME="happened-$BRANCH_NAME"
gcloud run services delete happened "$NAME" --region us-west1 --quiet