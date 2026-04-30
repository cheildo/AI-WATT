#!/usr/bin/env bash
# ── deploy-backend.sh ────────────────────────────────────────────────────────
# Generic SSH deploy helper for the aiwatt backend.
# Usage: ENV=staging ./scripts/deploy-backend.sh
# ─────────────────────────────────────────────────────────────────────────────
set -euo pipefail

ENV="${ENV:-staging}"
IMAGE="ghcr.io/neurowatt/aiwatt-backend"

case "$ENV" in
  staging)
    SSH_HOST="${STAGING_SSH_HOST:?STAGING_SSH_HOST not set}"
    SSH_USER="${STAGING_SSH_USER:?STAGING_SSH_USER not set}"
    SSH_KEY_FILE="${STAGING_SSH_KEY_FILE:-$HOME/.ssh/aiwatt_staging}"
    IMAGE_TAG="staging-latest"
    ENV_FILE="/opt/aiwatt/staging.env"
    ;;
  production)
    SSH_HOST="${PROD_SSH_HOST:?PROD_SSH_HOST not set}"
    SSH_USER="${PROD_SSH_USER:?PROD_SSH_USER not set}"
    SSH_KEY_FILE="${PROD_SSH_KEY_FILE:-$HOME/.ssh/aiwatt_prod}"
    IMAGE_TAG="latest"
    ENV_FILE="/opt/aiwatt/production.env"
    ;;
  *)
    echo "Unknown ENV: $ENV  (use staging or production)" >&2
    exit 1
    ;;
esac

SSH="ssh -i $SSH_KEY_FILE -o StrictHostKeyChecking=no $SSH_USER@$SSH_HOST"

echo "==> Deploying $IMAGE:$IMAGE_TAG to $ENV ($SSH_HOST)"

$SSH bash -s <<EOF
  set -euo pipefail
  docker pull ${IMAGE}:${IMAGE_TAG}
  docker stop aiwatt-api 2>/dev/null || true
  docker rm   aiwatt-api 2>/dev/null || true
  docker run -d \
    --name aiwatt-api \
    --restart unless-stopped \
    --env-file ${ENV_FILE} \
    -p 8080:8080 \
    ${IMAGE}:${IMAGE_TAG}
  docker image prune -f
  echo "Container started: \$(docker ps --filter name=aiwatt-api --format '{{.Status}}')"
EOF

echo "==> Deploy complete"
