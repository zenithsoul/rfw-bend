package main

import (
	"fmt"
	fwPaloAlto "rfw-bend/firewall/PaloAlto"
	modelGet "rfw-bend/model"
)

func main() {

	fmt.Println(fwPaloAlto.GetZonePalo())
	fmt.Println(modelGet.GetTopic())
}
