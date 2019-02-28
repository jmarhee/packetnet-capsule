variable "auth_token" {
	description = "Your Packet API Token"
}
variable "packet_ro_token" {
	description = "Your Read-Only Packet API Token (Will be passed to host)"
}
variable "packet_seek_tag" {
	description = "The tag capsule will filter nodes for (can be an empty string to make project-wide)"
	default = "capsule"
}
variable "packet_public_network" {
	description = "Firewall public network access (limits to hosts in project/tag -- accessible from jumphost"
	default = "true"
}

provider "packet" {
  auth_token = "${var.auth_token}"
}

resource "packet_project" "packetnet_capsule" {
  name           = "Packet Network"
}

resource "packet_device" "jumphost" {
  hostname         = "jump-capsule"
  plan             = "baremetal_0"
  facility         = "ewr1"
  operating_system = "ubuntu_18_04"
  billing_cycle    = "hourly"
  project_id       = "${packet_project.packetnet_capsule.id}"
  tags = ["capsule"]
}

data "template_file" "node" {
  template = "${file("${path.module}/node.tpl")}"

  vars {
    packet_ro_token = "${var.packet_ro_token}"
    packet_project_id = "${packet_project.packetnet_capsule.id}"
    packet_seek_tag    = "${var.packet_seek_tag}"
    packet_public_network = "${var.packet_public_network}"
  }
}

resource "packet_device" "capsule-host" {
  hostname         = "${format("capsule-test-node-%02d", count.index)}"
  count = "2"
  plan             = "baremetal_0"
  facility         = "ewr1"
  operating_system = "ubuntu_18_04"
  billing_cycle    = "hourly"
  project_id       = "${packet_project.packetnet_capsule.id}"
  tags = ["capsule"]
  user_data        = "${data.template_file.node.rendered}"
}


