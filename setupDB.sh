#!/bin/bash

set -e 

DB_USER="postgres"
DB_PASSWORD="postgres"

# Install PostgreSQL with enabled module
sudo dnf module list postgresql 
sudo dnf module enable -y postgresql:16
sudo dnf install -y postgresql-server 

# Initialize PostgreSQL database
sudo postgresql-setup --initdb

# Start PostgreSQL service
sudo systemctl start postgresql

# Enable PostgreSQL service to start on boot
sudo systemctl enable postgresql

# Change the password for the PostgreSQL default user 'postgres'
sudo -u postgres psql -c "ALTER USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"

# Update PostgreSQL to listen on all addresses
echo "listen_addresses = '*'" | sudo tee -a /var/lib/pgsql/data/postgresql.conf

# Allow incoming connections to PostgreSQL from all hosts
echo "host all all 0.0.0.0/0  md5" | sudo tee -a /var/lib/pgsql/data/pg_hba.conf

# Reload PostgreSQL for the changes to take effect
sudo systemctl reload postgresql