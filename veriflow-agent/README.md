# Veriflow Agent

Standalone Go binary deployed on borrower hardware as a systemd service.

Reads GPU and system health metrics every 5 minutes, signs the payload with HMAC-SHA256,
and POSTs it to the AI WATT backend API for ingestion and scoring.

**Reads:** `nvidia-smi` (GPU utilization, temperature, memory, power draw)  
**Reads:** `ipmitool` (fan speed — best-effort, non-fatal if missing)  
**Posts to:** `POST /api/v1/veriflow/telemetry`

---

## Prerequisites

- Go 1.22+
- `nvidia-smi` available on PATH (NVIDIA driver installed)
- `ipmitool` available on PATH (optional — for fan speed)

---

## Build

```bash
# Static binary — no CGO, single file
CGO_ENABLED=0 go build -o veriflow-agent ./cmd

# Cross-compile for Linux from macOS
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o veriflow-agent-linux ./cmd
```

---

## Configuration

```bash
cp .env.example .env
# Fill in ASSET_ID, HMAC_KEY, BACKEND_URL
```

| Variable | Description |
|---|---|
| `ASSET_ID` | UUID of the hardware asset — assigned during onboarding |
| `HMAC_KEY` | Shared secret for HMAC-SHA256 signing (32+ random bytes) |
| `BACKEND_URL` | Base URL of the AI WATT API (e.g. `https://api.aiwatt.io`) |

---

## Run

```bash
# Load env and run
export $(cat .env | xargs) && ./veriflow-agent
```

---

## systemd service

Create `/etc/systemd/system/veriflow-agent.service`:

```ini
[Unit]
Description=Veriflow Telemetry Agent
After=network.target

[Service]
Type=simple
User=veriflow
EnvironmentFile=/etc/veriflow-agent/env
ExecStart=/usr/local/bin/veriflow-agent
Restart=on-failure
RestartSec=30s

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable veriflow-agent
sudo systemctl start veriflow-agent
sudo journalctl -u veriflow-agent -f
```
