package main

import (
	"fmt"
	"github.com/OpenIoTHub/getip/iputils"
)

func main() {
	ipv4, err := iputils.GetMyPublicIpv4()
	if err != nil {
		fmt.Println("got my public ipv4 addr:", ipv4)
	}
	ipv6, err := iputils.GetMyPublicIpv6()
	if err != nil {
		fmt.Println("got my public ipv6 addr:", ipv6)
	}
}
