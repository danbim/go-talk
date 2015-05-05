package ipaddr

import "fmt"

type IPAddr [4]byte

var LOCALHOST = IPAddr{127, 0, 0, 1}

func (addr IPAddr) String() string {
	if addr == LOCALHOST {
		return "localhost"
	}
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
