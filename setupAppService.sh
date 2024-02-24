#!/bin/bash

# Configure systemd service to start on boot
sudo systemctl daemon-reload
sudo systemctl enable webapp.service

sudo systemctl start webapp.service

sudo systemctl status webapp.service
# log output from the service
journalctl -u webapp.service