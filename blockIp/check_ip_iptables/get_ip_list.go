package check_ip_iptables

import (
//	"fmt"
	"io/ioutil"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetIpList() []string {
	// 保存IP
	slice := []string{}
	// 读取文件
	file := "/var/log/maillog"
	dat, err := ioutil.ReadFile(file)
	check(err)
	// fmt.Print(string(dat))
	source := string(dat)
	//匹配IP
	pattern := `((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`
	reg := regexp.MustCompile(pattern)
	ip := reg.FindAllString(source, -1)
	// fmt.Printf("%s", reg.FindAllString(source, -1))
	// fmt.Printf("%s", ip)
	for _, v := range ip {
		// fmt.Println(v)
		slice = append(slice, v)
	}
	// 返回匹配到的IP
	m := RmDuplicate(&slice)
	return m
}

// slice去重
func RmDuplicate(list *[]string) []string {
	var x []string = []string{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
