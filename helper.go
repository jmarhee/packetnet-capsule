package main

import (
	"strings"
	"net"
	"errors"
)

// FindInterfaceName returns the network interface name of the provided local
// ip address
func FindInterfaceName(ifaces []net.Interface, local string) (string, error) {
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPAddr:
				ip := v.IP.String()
				if ip == local {
					return i.Name, nil
				}
			case *net.IPNet:
				ip := v.IP.String()
				if ip == local {
					return i.Name, nil
				}
			}
		}
	}

	return "", errors.New("local interface could not be found")
}

func PullIPAddress(targetType string) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			if strings.HasPrefix(ip.String(), "10.") == true {
				if targetType == "private" {
					return ip.String(), nil
				} else {
					continue
				}
			} else {
				if targetType == "public" {
					return ip.String(), nil
				} else {
					continue
				}
			}
		}
	}

	return "", errors.New("are you connected to the network?")
}

func PrivateAddress() (string, error) {
	privateIp, err := PullIPAddress("private")
	if err != nil {
                return "", errors.New("Cannot detect IP.")
	}

	return privateIp, nil
}

func PublicAddress() (string, error) {
        publicIp, err := PullIPAddress("public")
        if err != nil {
                return "", errors.New("Cannot detect IP.")
        }

        return publicIp, nil
}
