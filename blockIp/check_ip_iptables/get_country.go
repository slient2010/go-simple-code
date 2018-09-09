package check_ip_iptables

import (
	"fmt"
	"github.com/abh/geoip"
)

func GetIPCountry(ip string) string {
	file := "/usr/share/GeoIP/GeoIP.dat"

	gi, err := geoip.Open(file)
	if err != nil {
		fmt.Printf("Could not open GeoIP database\n")
	}

	// country, _ := gi.GetCountry("58.250.251.47")
	country, _ := gi.GetCountry(ip)
	return country
}
