#!/bin/bash

# Configure systemd service to start on boot
sudo systemctl daemon-reload
sudo systemctl enable webapp.service

# log output from the service
journalctl -u webapp.service