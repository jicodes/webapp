#!/bin/bash

set -e 

DIR="/opt/myapp"

sudo mkdir -p "${DIR}" # Create the directory

sudo chown -R packer:packer "${DIR}" # Change the owner of the directory to 'packer'