#!/bin/bash

# Configure systemd service to start on boot
sudo systemctl daemon-reload
sudo systemctl enable webapp.service

# Start the service
sudo systemctl start webapp.service

# Check the status of the service
sudo systemctl status webapp.service