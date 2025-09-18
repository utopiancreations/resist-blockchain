#!/bin/bash
set -e
apt-get update
apt-get install -y build-essential
curl https://get.ignite.com/cli | bash
