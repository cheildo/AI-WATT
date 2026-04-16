BINARY   := veriflow-agent-linux-amd64
DIST_DIR := dist
CMD_DIR  := ./cmd

.PHONY: build-agent clean install

## build-agent: compile a static Linux amd64 binary into dist/
build-agent:
	@mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -ldflags="-s -w" -o $(DIST_DIR)/$(BINARY) $(CMD_DIR)
	@echo "Built: $(DIST_DIR)/$(BINARY)"

## install: copy binary to /usr/local/bin and reload systemd unit
install: build-agent
	install -m 0755 $(DIST_DIR)/$(BINARY) /usr/local/bin/veriflow-agent
	install -m 0644 deploy/veriflow-agent.service /etc/systemd/system/veriflow-agent.service
	systemctl daemon-reload
	systemctl enable --now veriflow-agent

## clean: remove build artifacts
clean:
	rm -rf $(DIST_DIR)
