package main

import (
	"log"
	"github.com/packethost/packngo"
)

func SeekTag(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func DeviceList(projectId string, ifaceType int, tag string) []string {

	deviceAddresses := []string{}

	c, err := packngo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ds, _, err := c.Devices.List(projectId, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range ds {
		if SeekTag(tag, d.Tags) || len(tag) == 0 {
			deviceAddresses = append(deviceAddresses, string(d.Network[ifaceType].Address))
		}
	}

	return deviceAddresses
}

func PrivateDevices(projectId string, tag string) []string {
	devicePrivateAddresses := DeviceList(projectId, 2, tag)

	return devicePrivateAddresses
}

func PublicDevices(projectId string, tag string) []string {
        devicePublicAddresses := DeviceList(projectId, 0, tag)

        return devicePublicAddresses
}
