#!/bin/bash

# Build the application
go build -o space-monitor ./cmd/space-monitor

# Copy binary to /usr/local/bin
sudo cp space-monitor /usr/local/bin/

# Copy systemd service file
sudo cp systemd/space-monitor.service /etc/systemd/system/

# Reload systemd and start service
sudo systemctl daemon-reload
sudo systemctl enable space-monitor
sudo systemctl start space-monitor

# Run database migrations
psql -U postgres -d space_monitor -f migrations/0001_init.sql
