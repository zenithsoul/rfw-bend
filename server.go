package main

import (
	"rfw-bend/firewall"
)

func main() {
	firewall.InitPalo()

	c := firewall.GetZonePalo()
}
