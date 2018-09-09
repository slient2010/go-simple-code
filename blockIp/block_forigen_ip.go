package main

import (
	. "check_ip_iptables"
	"time"
	"fmt"
)

func main() {
	ipt, err := New()
	if err != nil {
		fmt.Printf("New failed: %v", err)
	}
	for _, v := range GetIpList() {
		if GetIPCountry(v) != "CN" && GetIPCountry(v) != "" {
			fmt.Println(v)
			err = ipt.AppendUnique("filter", "INPUT", "-s", v, "-j", "DROP")
			if err != nil {
				fmt.Printf("AppendUnique failed: %v", err)
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}
}
