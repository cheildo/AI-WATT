#!/usr/bin/env bash
# ── health-check.sh ──────────────────────────────────────────────────────────
# Post-deploy health verification for API and frontend.
# Usage: ENV=staging ./scripts/health-check.sh
# ─────────────────────────────────────────────────────────────────────────────
set -euo pipefail

ENV="${ENV:-staging}"
MAX_RETRIES="${MAX_RETRIES:-10}"
RETRY_DELAY="${RETRY_DELAY:-10}"

case "$ENV" in
  staging)
    API_URL="https://api-staging.aiwatt.io"
    FRONTEND_URL="https://staging.aiwatt.io"
    ;;
  production)
    API_URL="https://api.aiwatt.io"
    FRONTEND_URL="https://aiwatt.io"
    ;;
  local)
    API_URL="http://localhost:8080"
    FRONTEND_URL="http://localhost:5173"
    ;;
  *)
    echo "Unknown ENV: $ENV  (use local, staging or production)" >&2
    exit 1
    ;;
esac

# ── API health check ─────────────────────────────────────────────────────────
echo "==> Checking API: $API_URL/api/v1/vault/stats"
for i in $(seq 1 "$MAX_RETRIES"); do
  STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/api/v1/vault/stats" 2>/dev/null || echo "000")
  echo "    Attempt $i/$MAX_RETRIES: HTTP $STATUS"
  if [ "$STATUS" = "200" ]; then
    echo "    API is healthy"
    break
  fi
  if [ "$i" -eq "$MAX_RETRIES" ]; then
    echo "ERROR: API health check failed after $MAX_RETRIES attempts" >&2
    exit 1
  fi
  sleep "$RETRY_DELAY"
done

# ── Frontend health check ─────────────────────────────────────────────────────
echo "==> Checking frontend: $FRONTEND_URL"
STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$FRONTEND_URL" 2>/dev/null || echo "000")
echo "    HTTP $STATUS"
if [ "$STATUS" != "200" ]; then
  echo "ERROR: Frontend returned HTTP $STATUS" >&2
  exit 1
fi
echo "    Frontend is healthy"

echo "==> All health checks passed for $ENV"
