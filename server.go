package main

import (
	"fmt"
	firewall "fwapi"
)

func main() {
	firewall.InitPalo()

	c, _ := firewall.GetZonePalo()

	fmt.Println(c)
}
