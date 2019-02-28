packetnet-capsule
===

Inspired by [droplan](https://github.com/tam7t/droplan), this package configures your [Packet](https://packet.com) host firewalls to limit traffic only to those hosts--either project-wide, or to a tag-based subset of hosts running the agent, which will periodically update from the Packet API.

Ideal uses might be highly-network dependent frameworks like [Kubernetes on Packet](https://github.com/jmarhee/packet-multiarch-k8s-terraform), which can be highly automated, and requires only limited network access between hosts.

Running this package (i.e. via cronjob, or in a service) will update rules at regular intervals to ensure the rules are kept current with your specification (either project-wide, or limited to nodes with a hosts tag ID). 

Setup
---

This package is supported on `arm64` and `amd64` servers. 

`packetnet-capsule` requires 3 configuration variables:

`PACKET_AUTH_TOKEN`: a [read-only access key](https://www.packet.com/developers/changelog/project-only-api-keys/) in order to get data from the Packet API to keep firewalls up-to-date.

`PACKET_PROJECT_ID`: the project the hosts will reside in.

`PACKET_SEEK_TAG` (Optional): if set, hosts tagged with this value will be targetted.

`PUBLIC` (Optional): if set, will include Public IP addresses in ruleset (meaning these hosts will be inaccessible, except from other hosts in the network). 

Usage
---

The package can be built for `arm64` and `amd64` hosts using the Makefile:

```
make build
make build-arm
```

and then run:

```
PACKET_AUTH_TOKEN=<ro token> \
PACKET_SEEK_TAG="capsule" \
PACKET_PROJECT_ID=<id> ./packetnet-capsule
```

or using the Docker images:

```
jmarhee/packetnet-capsule.amd64
jmarhee/packetnet-capsule.arm64
```
as in:

```
docker run -d --restart=always --net=host --cap-add=NET_ADMIN \
--name packetnet-capsule \
-e PACKET_AUTH_TOKEN=$PACKET_AUTH_TOKEN \
-e PACKET_PROJECT_ID=$PACKET_PROJECT_ID \
-e PACKET_SEEK_TAG=$PACKET_SEEK_TAG \
-e PUBLIC=$PUBLIC \
-e CHECK_INTERVAL=300 jmarhee/packetnet-capsule.amd64:latest
```

Example
---

In `example/`, you can use Terraform to spin-up a test environment.

```
export TF_VAR_auth_token=<your RW API key>
export TF_VAR_packet_ro_token=<your RO API key>
export TF_VAR_packet_public_network="true"

terraform apply
```

This example creates two hosts running `packetnet-capsule`, and a host that is tagged `capsule` to allow access to the firewalled hosts, which is are not accessible from outside of nodes within that tag.



