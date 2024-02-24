#!/bin/bash

set -e 

DIR="/opt/myapp"

sudo mkdir -p "${DIR}"

sudo chown -R packer:packer "${DIR}" 