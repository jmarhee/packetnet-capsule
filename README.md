packetnet-fw-agent
===

[![Build Status](https://cloud.drone.io/api/badges/packet-labs/packetnet-fw-agent/status.svg)](https://cloud.drone.io/packet-labs/packetnet-fw-agent)

Inspired by [droplan](https://github.com/tam7t/droplan).

This package configures your [Packet](https://packet.com) host firewalls to limit traffic only to those hosts.

Either project-wide, or to a tag-based subset of hosts running the agent, modes available.
It will periodically update lists from the Packet API.

Ideal uses might be highly-network dependent frameworks like:

- [Kubernetes on Packet](https://github.com/jmarhee/packet-multiarch-k8s-terraform)

Cronjobs can be used to update rules dynamically.

Setup
---

This package is supported on `arm64` and `amd64` servers.

`packetnet-fw-agent` requires 3 configuration variables:

`PACKET_AUTH_TOKEN`: [read-only key](https://www.packet.com/developers/changelog/project-only-api-keys/)

`PACKET_PROJECT_ID`: the project the hosts will reside in.

`PACKET_SEEK_TAG` (Optional): if set, hosts tagged with this value will be targetted.

`PUBLIC` (Optional): if set, will include Public IP addresses in ruleset.
These hosts will be inaccessible, except from other hosts in the network).

Usage
---

The package can be built for `arm64` and `amd64` hosts using the Makefile:

```bash
make build
make build-arm
```

and then run:

```bash
PACKET_AUTH_TOKEN=<ro token> \
PACKET_SEEK_TAG="capsule" \
PACKET_PROJECT_ID=<id> ./packetnet-fw-agent
```

or using the Docker images (on Docker Hub):

[jmarhee/packetnet-fw-agent.amd64](https://cloud.docker.com/repository/docker/jmarhee/packetnet-fw-agent.amd64)

[jmarhee/packetnet-fw-agent.arm64](https://cloud.docker.com/repository/docker/jmarhee/packetnet-fw-agent.arm64)

as in:

```bash
docker run -d --restart=always --net=host --cap-add=NET_ADMIN \
--name packetnet-fw-agent \
-e PACKET_AUTH_TOKEN=$PACKET_AUTH_TOKEN \
-e PACKET_PROJECT_ID=$PACKET_PROJECT_ID \
-e PACKET_SEEK_TAG=$PACKET_SEEK_TAG \
-e PUBLIC=$PUBLIC \
-e CHECK_INTERVAL=300 jmarhee/packetnet-fw-agent.amd64:latest
```

You can build the Docker images using the Makefile:

```bash
make TAG=$(date +%F%H%M%S) docker-arm64
make TAG=$(date +%F%H%M%S) docker-amd64
```

which will build new binaries as well for the desired architecture.

Example
---

In `example/`, you can use Terraform to spin-up a test environment.

```bash
export TF_VAR_auth_token=<your RW API key>
export TF_VAR_packet_ro_token=<your RO API key>
export TF_VAR_packet_public_network="true"

terraform apply
```

This example creates two hosts running `packetnet-fw-agent`.
The tagged node will be inaccessible from outside the network.
