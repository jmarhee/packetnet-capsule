#!/bin/bash

export PACKET_AUTH_TOKEN=${packet_ro_token}
export PACKET_PROJECT_ID=${packet_project_id}
export PACKET_SEEK_TAG=${packet_seek_tag}
export PUBLIC=${packet_public_network}

apt update; \
apt install -y docker.io

docker run -d --restart=always --net=host --cap-add=NET_ADMIN \
-e PACKET_AUTH_TOKEN=$PACKET_AUTH_TOKEN \
-e PACKET_PROJECT_ID=$PACKET_PROJECT_ID \
-e PACKET_SEEK_TAG=$PACKET_SEEK_TAG \
-e PUBLIC=$PUBLIC \
-e CHECK_INTERVAL=300 jmarhee/packetnet-capsule.amd64:latest
