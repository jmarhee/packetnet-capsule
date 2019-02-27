package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/coreos/go-iptables/iptables"
)

var appVersion string

func main() {
	version := flag.Bool("version", false, "Print the version and exit.")
	flag.Parse()
	if *version {
		log.Printf(appVersion)
		os.Exit(0)
	}

	accessToken := os.Getenv("PACKET_AUTH_TOKEN")
	if accessToken == "" {
		log.Fatal("Usage: PACKET_AUTH_TOKEN environment variable must be set.")
	}

	// PUBLIC=true will tell us to block traffic on the public interface
	public := os.Getenv("PUBLIC")

	// PACKET_PROJECT_ID
	projectId := os.Getenv("PACKET_PROJECT_ID")

	// PACKET_SEEK_TAG=<string> can be used to limit to devices with this tag
	tag := os.Getenv("PACKET_SEEK_TAG")

	// setup dependencies
	ipt, err := iptables.New()
	failIfErr(err)

	// collect local network interface information
	ifaces, err := net.Interfaces()
	failIfErr(err)

	pubAddr, err := PublicAddress()
	failIfErr(err)

	if public == "true" {
		publicPeers := PublicDevices(projectId, tag)

		// find public iface name
		iface, err := FindInterfaceName(ifaces, pubAddr)
		failIfErr(err)

		// setup packetnet-peers-public chain for public interface
		err = Setup(ipt, iface, "packetnet-peers-public")
		failIfErr(err)

		// update packetnet-peers-public
		err = UpdatePeers(ipt, publicPeers, "packetnet-peers-public")
		failIfErr(err)
		log.Printf("Added %d peers to packetnet-peers-public", len(publicPeers))
	}

	privAddr, err := PrivateAddress()
	failIfErr(err)

	// find private iface name
	iface, err := FindInterfaceName(ifaces, privAddr)
	if public != "" && err != nil && err.Error() == "no private interfaces" {
		os.Exit(0)
	}
	failIfErr(err)

	// setup packetnet-peers chain for private interface
	err = Setup(ipt, iface, "packetnet-peers")
	failIfErr(err)

	privatePeers := PrivateDevices(projectId, tag)

	// update packetnet-peers
	err = UpdatePeers(ipt, privatePeers, "packetnet-peers")
	failIfErr(err)
	log.Printf("Added %d peers to packetnet-peers", len(privatePeers))
}

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
