package ipaddr

import (
	"testing"
)

type TestIPAddr struct {
	address  IPAddr
	expected string
}

var tests = []TestIPAddr{
	{IPAddr{127, 0, 0, 1}, "localhost"},
	{IPAddr{192, 168, 0, 1}, "192.168.0.1"},
	{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	{IPAddr{255, 255, 255, 0}, "255.255.255.0"},
}

func TestStringer(t *testing.T) {
	for _, ip := range tests {
		actual := ip.address.String()
		if ip.expected != actual {
			t.Errorf("Got %v, expected %v", actual, ip.expected)
		}
	}
}

func BenchmarkIPAddr(b *testing.B) {
	for _, ip := range tests {
		for i := 0; i < b.N; i++ {
			ip.address.String()
		}
	}
}
