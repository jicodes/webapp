#!/bin/bash

set -e  # Exit on error
set -o pipefail  # Exit on pipeline command failure

# Install PostgreSQL with enabled module
sudo dnf module list postgresql 
sudo dnf module enable -y postgresql:16
sudo dnf install -y postgresql-server 

echo "Initializing PostgreSQL database..."
sudo postgresql-setup --initdb

# Start PostgreSQL service
sudo systemctl start postgresql

# Enable PostgreSQL service to start on boot
sudo systemctl enable postgresql

# Change the password for the PostgreSQL default user 'postgres'
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'postgres';"

# Update PostgreSQL to listen on all addresses
echo "listen_addresses = '*'" | sudo tee -a /var/lib/pgsql/data/postgresql.conf

# Allow incoming connections to PostgreSQL from all hosts
echo "host all all 0.0.0.0/0  md5" | sudo tee -a /var/lib/pgsql/data/pg_hba.conf
echo "host all all ::0/0  md5" | sudo tee -a /var/lib/pgsql/data/pg_hba.conf

# Allow local connections to PostgreSQL from the postgres user
echo "local all all md5" | sudo tee -a /var/lib/pgsql/data/pg_hba.conf

# Restart PostgreSQL for the changes to take effect
echo "Starting PostgreSQL service..."
sudo systemctl restart postgresql