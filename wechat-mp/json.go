// 解析json的例子
package main

import "fmt"
import "encoding/json"

type MxRecords struct {
	Value    string
	Ttl      int
	Priority int
	HostName string
}

type Data struct {
	MxRecords []MxRecords
}

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
type apiR struct {
	Response Response `json:"response"`
}

func main() {

	body := `
  {"response": {
  "status": "SUCCESS",
  "data": {
    "mxRecords": [
      {
        "value": "us2.mx3.mailhostbox.com.",
        "ttl": 1,
        "priority": 100,
        "hostName": "hostname@"
      },
      {
        "value": "us2.mx1.mailhostbox.com.",
        "ttl": 1,
        "priority": 100,
        "hostName": "hostnam233e@"
      },
      {
        "value": "us2.mx2.mailhostbox.com.",
        "ttl": 1,
        "priority": 100,
        "hostName": "@"
      }
    ],
    "cnameRecords": [
      {
        "aliasHost": "pop.a.co.uk.",
        "canonicalHost": "us2.pop.mailhostbox.com."
      },
      {
        "aliasHost": "webmail.a.co.uk.",
        "canonicalHost": "us2.webmail.mailhostbox.com."
      },
      {
        "aliasHost": "smtp.a.co.uk.",
        "canonicalHost": "us2.smtp.mailhostbox.com."
      },
      {
        "aliasHost": "imap.a.co.uk.",
        "canonicalHost": "us2.imap.mailhostbox.com."
      }
    ],
    "dkimTxtRecord": {
      "domainname": "20a19._domainkey.a.co.uk",
      "value": "\"v=DKIM1; g=*; k=rsa; p=DkfbhO8Oyy0E1WyUWwIDAQAB\"",
      "ttl": 1
    },
    "spfTxtRecord": {
      "domainname": "a.co.uk",
      "value": "\"v=spf1 redirect=_spf.mailhostbox.com\"",
      "ttl": 1
    },
    "loginUrl": "us2.cp.mailhostbox.com"
  }
}}`

	var r apiR
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Println(r)
	fmt.Println(r.Response.Data.MxRecords[1].HostName)

}
