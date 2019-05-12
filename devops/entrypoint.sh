#!/bin/sh
if [ $DOCKER_HOST ]; then
  echo "Docker Host: $DOCKER_HOST (manual override)"
else
  DOCKER_HOST="$(getent hosts host.docker.internal | cut -d' ' -f1)"
  if [ $DOCKER_HOST ]; then
    echo "Docker Host: $DOCKER_HOST (host.docker.internal)"
  else
    DOCKER_HOST=$(ip -4 route show default | cut -d' ' -f3)
    echo "Docker Host: $DOCKER_HOST (default gateway)"
  fi
fi

echo "$DOCKER_HOST host.docker.internal" >> /etc/hosts

./content-api