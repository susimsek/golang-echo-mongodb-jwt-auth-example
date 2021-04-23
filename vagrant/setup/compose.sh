#!/bin/bash

export DOCKER_COMPOSE_VERSION=1.27.4

# Introduced in 1-4.3
sudo curl -L https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m) -o /usr/bin/docker-compose
 
sudo chmod +x /usr/bin/docker-compose

echo "======= Done. PLEASE LOG OUT & LOG Back In ===="
echo "Then validate by executing    'docker-compose version'"