package main

import (
	"fmt"
	"net"
	"net/http"
	// "strings"
	"encoding/json"
)

type IPaddress struct {
	Ip string `json:"ip"`
}

func GetIP(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// fmt.Fprintln(return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
		fmt.Fprintln(w, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr))
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		// return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
		fmt.Fprintln(w, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr))
	}
	var userip IPaddress
	userip.Ip = userIP.String()
	data, e := json.Marshal(userip)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Fprintln(w, string(data))

}
