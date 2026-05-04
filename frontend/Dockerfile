## ── Build stage ──────────────────────────────────────────────────────────────
FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci --prefer-offline

COPY . .

# VITE_* vars are baked into the static bundle at build time via docker-compose args
ARG VITE_ENV=production
ARG VITE_CHAIN_ID
ARG VITE_RPC_URL
ARG VITE_API_URL
ARG VITE_WALLETCONNECT_PROJECT_ID
ARG VITE_WATT_ADDRESS
ARG VITE_SWATT_ADDRESS
ARG VITE_MINT_ENGINE_ADDRESS
ARG VITE_LENDING_POOL_ADDRESS
ARG VITE_WEV_QUEUE_ADDRESS
ARG VITE_USDC_ADDRESS
ARG VITE_USDT_ADDRESS

ENV VITE_ENV=$VITE_ENV \
    VITE_CHAIN_ID=$VITE_CHAIN_ID \
    VITE_RPC_URL=$VITE_RPC_URL \
    VITE_API_URL=$VITE_API_URL \
    VITE_WALLETCONNECT_PROJECT_ID=$VITE_WALLETCONNECT_PROJECT_ID \
    VITE_WATT_ADDRESS=$VITE_WATT_ADDRESS \
    VITE_SWATT_ADDRESS=$VITE_SWATT_ADDRESS \
    VITE_MINT_ENGINE_ADDRESS=$VITE_MINT_ENGINE_ADDRESS \
    VITE_LENDING_POOL_ADDRESS=$VITE_LENDING_POOL_ADDRESS \
    VITE_WEV_QUEUE_ADDRESS=$VITE_WEV_QUEUE_ADDRESS \
    VITE_USDC_ADDRESS=$VITE_USDC_ADDRESS \
    VITE_USDT_ADDRESS=$VITE_USDT_ADDRESS

RUN npm run build

## ── Runtime stage ────────────────────────────────────────────────────────────
FROM nginx:1.27-alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
