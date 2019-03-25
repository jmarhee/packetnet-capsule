#!/bin/sh

set -ex

if [ -z "$CHECK_INTERVAL" ]; then
	# if the interval is not set, only execute once
	PACKET_AUTH_TOKEN=$PACKET_AUTH_TOKEN PACKET_PROJECT_ID=$PACKET_PROJECT_ID PACKET_SEEK_TAG=$PACKET_SEEK_TAG ./packetnet-fw-agent
else
	while true; do
		# since we use 'set -e', this while loop will exit if droplan exits with a return value other than 0
		# (which in turn tells docker to restart the container (assuming the --restart option was used)
		# while delaying retries exponentially)
		PACKET_AUTH_TOKEN=$PACKET_AUTH_TOKEN PACKET_PROJECT_ID=$PACKET_PROJECT_ID PACKET_SEEK_TAG=$PACKET_SEEK_TAG ./packetnet-fw-agent
		sleep "$CHECK_INTERVAL"
	done
fi
