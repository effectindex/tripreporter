#!/bin/sh

NAME="tripreporter"

if [ -z "$(which docker)" ]; then
  echo "docker not found in path! exiting"
  exit 1
fi

# shellcheck disable=SC2039
if ! source "$PWD/.env"; then
  echo "failed to \`source .env\`, does it exist? exiting"
  exit 1
fi

if [ -z "$SRV_PORT" ]; then
  echo "SRV_PORT not set in \`.env\`, exiting"
  exit 1
fi

docker build . -t "$NAME" || exit 1

if [ -n "$(docker ps --filter name="$NAME" --format "{{.ID}}" | head -n 1)" ]; then
  docker stop "$NAME" || {
    echo "failed to stop old container, is it running?"
  }
fi

if [ -n "$(docker container ls -a --filter name="$NAME" --format "{{.ID}}" | head -n 1)" ]; then
  docker rm "$NAME" || {
    echo "failed to remove old container"
  }
fi

docker run -d -p "127.0.0.1:$SRV_PORT:$SRV_PORT" --restart=unless-stopped --name "$NAME" "$NAME"
